package confluence

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"smart_confluence_search/constant"
	"smart_confluence_search/lib"

	"strings"
)

func GetPageContentByID(pageID string) PageContent {
	confluenceEndpoint := os.Getenv("CONFLUENCE_ENDPOINT")
	url := confluenceEndpoint + fmt.Sprintf(constant.GetPageContenAPI, pageID) + "?body-format=view"

	req := getClient(url, "GET", nil)

	res, err := http.DefaultClient.Do(req)
	lib.HandleErr(err)
	defer res.Body.Close()

	var PageContent PageContent
	err = json.NewDecoder(res.Body).Decode(&PageContent)
	if err != nil {
		panic(err)
	}

	return PageContent
}

func UpdateTagsToPage(pageID string, tagsString string) *http.Response {
	confluenceEndpoint := os.Getenv("CONFLUENCE_ENDPOINT")
	url := confluenceEndpoint + fmt.Sprintf(constant.PostPageLabelAPI, pageID)

	input := Label{
		Prefix: "global",
		Name:   tagsString,
	}

	jsonData, err := json.Marshal(input)
	lib.HandleErr(err)

	req := getClient(url, "POST", bytes.NewBuffer(jsonData))

	res, err := http.DefaultClient.Do(req)
	lib.HandleErr(err)

	return res
}

func convertStringToSlice(str string) []string {
	str = strings.TrimPrefix(str, "[")
	str = strings.TrimSuffix(str, "]")
	str = strings.ReplaceAll(str, " ", "")
	return strings.Split(str, ",")
}

func SearchPageByLabel(labels string) []Result {
	url := genURL(convertStringToSlice(labels))
	fmt.Println("search url:", url)

	req := getClient(url, "GET", nil)

	res, err := http.DefaultClient.Do(req)
	lib.HandleErr(err)
	defer res.Body.Close()

	var SearchResult SearchResult
	err = json.NewDecoder(res.Body).Decode(&SearchResult)
	if err != nil {
		panic(err)
	}
	fmt.Println("search result:", SearchResult.Results)

	return SearchResult.Results
}

func getClient(url string, method string, body io.Reader) *http.Request {
	username := os.Getenv("CONFLUENCE_USERNAME")
	password := os.Getenv("CONFLUENCE_PASSWORD")

	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	return req
}

func genURL(labels []string) string {
	// cql=label in (label1, label2)
	// label%20in%28"label1"%2C"label2"%29
	labelLimit := labels[:7]
	searchLabels := strings.Join(labelLimit, "%2C")
	cql := "label%20in%28" + searchLabels[:len(searchLabels)-3] + "%29"

	confluenceEndpoint := os.Getenv("CONFLUENCE_ENDPOINT")
	url := confluenceEndpoint + constant.ConfluenceSearchAPI
	url += "space=" + constant.SpaceID
	url += "&limit=" + constant.ConfluenceSearchLimit
	url += "&cql=" + cql

	return url
}
