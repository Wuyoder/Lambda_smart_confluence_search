package main

import (
	"context"
	"fmt"
	"os"
	"smart_confluence_search/confluence"
	"smart_confluence_search/constant"
	"smart_confluence_search/lib"
	"smart_confluence_search/openai"

	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {
	lambda.Start(Handler)
}

// GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap -tags lambda.norpc main.go
// zip myFunction.zip bootstrap

type Search struct {
	Search string `json:"search"`
}

type SearchResultList struct {
	Message          string   `json:"message"`
	SearchResultList []string `json:"search_result_list"`
}

func Handler(event Search) (SearchResultList, error) {
	timeStart := time.Now()
	ctx := context.Background()
	fmt.Println("Start to search")
	var output SearchResultList
	searchString := event.Search
	if searchString == "" {
		fmt.Println("query is empty")
		output.Message = "query is empty"
		return output, nil
	}

	fmt.Println("start to get label page")
	labelPageID := os.Getenv("LABEL_PAGE_ID")
	labels := confluence.GetPageContentByID(labelPageID)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(labels.Body.View.Value))
	lib.HandleErr(err)
	labelList := []string{}
	doc.Find("a.label").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		labelList = append(labelList, text)
	})
	fmt.Println("labelList : ", labelList)

	getSearchCompletion := openai.CompletionObj{
		Payload:    searchString,
		Query:      fmt.Sprintf(constant.QuerySearch, strings.Join(labelList, ",")),
		IsChatMode: true,
		IsGPT4:     true,
	}
	var pageIDs []string
	fmt.Println("start to get search completion")
	searchList, _ := getSearchCompletion.GetOpenAIResp(ctx)
	fmt.Println("AI parse labels :", searchList[0])
	fmt.Println("start to search page by label")
	pages := confluence.SearchPageByLabel(searchList[0])

	confluenceEndpoint := os.Getenv("CONFLUENCE_ENDPOINT")

	if len(pages) == 0 {
		fmt.Println("No page found")
		output.Message = "No page found"
		lib.PostToSlack(lib.SlackWebhook, lib.SlackChannel, output.Message)
		return output, nil
	} else {
		output.Message = "Search result :"
		lib.PostToSlack(lib.SlackWebhook, lib.SlackChannel, output.Message)
		for _, v := range pages {
			fmt.Println("pageID : ", v.Content.ID)
			url := confluenceEndpoint + constant.SpacePagePath + v.Content.ID
			pageIDs = append(pageIDs, url)
			lib.PostToSlack(lib.SlackWebhook, lib.SlackChannel, url)
		}
		output.SearchResultList = pageIDs
	}
	endTime := time.Now()
	fmt.Println("Time taken:", endTime.Sub(timeStart))

	return output, nil
}

func init() {
	godotenv.Load()
}
