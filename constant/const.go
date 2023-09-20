package constant

const (
	QueryLabel  = "這是一份html內容，包含數個標籤，請將標籤內容提取出來，並用空格將標籤區隔，不需要說明與標點符號。"
	QuerySearch = "請將以下內容歸納理解，並將內容以提供的標籤集合作對比，從標籤集合物中，把與內容批配的標籤列出來。 提供的標籤集合內容如下:```%s```。提供的內容如下："
)

var (
	GetPageContenAPI      = "wiki/api/v2/pages/%s"
	PostPageLabelAPI      = "wiki/rest/api/content/%s/label"
	SpacePagePath         = "wiki/spaces/PS3/pages/"
	ConfluenceSearchAPI   = "wiki/rest/api/search?"
	SpaceID               = "PS3"
	ConfluenceSearchLimit = "5"
)
