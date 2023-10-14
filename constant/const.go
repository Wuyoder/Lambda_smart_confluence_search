package constant

import (
	"fmt"
	"os"
)

const (
	QuerySearch = "Please summarize and understand the following content, and compare it with the tag set provided. From the set of tags provided, list out the tags that match the content. The content of the tag set provided is as follows: ```%s```. The content provided is as follows:"
)

var (
	GetPageContenAPI      = "wiki/api/v2/pages/%s"
	PostPageLabelAPI      = "wiki/rest/api/content/%s/label"
	SpaceID               = os.Getenv("SPACE_ID")
	SpacePagePath         = fmt.Sprintf("wiki/spaces/%s/pages/", SpaceID)
	ConfluenceSearchAPI   = "wiki/rest/api/search?"
	ConfluenceSearchLimit = "5"
)
