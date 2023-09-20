package confluence

type PageContent struct {
	ParentType string    `json:"parentType"`
	AuthorID   string    `json:"authorId"`
	ID         string    `json:"id"`
	Version    Version   `json:"version"`
	Position   int       `json:"position"`
	Title      string    `json:"title"`
	Status     string    `json:"status"`
	Body       Body      `json:"body"`
	ParentID   string    `json:"parentId"`
	SpaceID    string    `json:"spaceId"`
	CreatedAt  string    `json:"createdAt"`
	Links      LinksPage `json:"_links"`
}

type Version struct {
	Number    int    `json:"number"`
	Message   string `json:"message"`
	MinorEdit bool   `json:"minorEdit"`
	AuthorID  string `json:"authorId"`
	CreatedAt string `json:"createdAt"`
}

type Body struct {
	View View
}

type View struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type LinksPage struct {
	EditUI string `json:"editui"`
	WebUI  string `json:"webui"`
	TinyUI string `json:"tinyui"`
}

type Label struct {
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
}

type SearchResult struct {
	Results        []Result    `json:"results"`
	Start          int         `json:"start"`
	Limit          int         `json:"limit"`
	Size           int         `json:"size"`
	TotalSize      int         `json:"totalSize"`
	CqlQuery       string      `json:"cqlQuery"`
	SearchDuration int         `json:"searchDuration"`
	Links          LinksSearch `json:"_links"`
}

type LinksSearch struct {
	Base    string `json:"base"`
	Context string `json:"context"`
	Self    string `json:"self"`
	Next    string `json:"next"`
}

type Result struct {
	Content               Content               `json:"content"`
	Title                 string                `json:"title"`
	Excerpt               string                `json:"excerpt"`
	Url                   string                `json:"url"`
	ResultBlobalContainer ResultBlobalContainer `json:"resultGlobalContainer"`
	BreadCrumb            interface{}           `json:"breadcrumb"`
	EntityType            string                `json:"entityType"`
	IconCssClass          string                `json:"iconCssClass"`
	LastModified          string                `json:"lastModified"`
	FriendlyLastModified  string                `json:"friendlyLastModified"`
	Score                 float64               `json:"score"`
}

type Content struct {
	ID                  string      `json:"id"`
	Type                string      `json:"type"`
	Status              string      `json:"status"`
	Title               string      `json:"title"`
	ChildTypes          interface{} `json:"childTypes"`
	MacroRenderedOutput interface{} `json:"macroRenderedOutput"`
	Restrictions        interface{} `json:"restrictions"`
	Expandable          Expandable  `json:"_expandable"`
	Links               LinksResult `json:"_links"`
}

type LinksResult struct {
	WebUI  string `json:"webui"`
	Self   string `json:"self"`
	TinyUI string `json:"tinyui"`
}

type Expandable struct {
	Container   string `json:"container"`
	Metadata    string `json:"metadata"`
	Extensions  string `json:"extensions"`
	Operations  string `json:"operations"`
	Children    string `json:"children"`
	History     string `json:"history"`
	Ancestors   string `json:"ancestors"`
	Body        string `json:"body"`
	Version     string `json:"version"`
	Descendants string `json:"descendants"`
	Space       string `json:"space"`
}

type ResultBlobalContainer struct {
	Title      string `json:"title"`
	DisplayUrl string `json:"displayUrl"`
}
