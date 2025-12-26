package controller

const (
	NEWS_LIMIT                 = 10
	NEWS_SEARCH_LIMIT          = 50
	NEWS_PAGE                  = 1
	CategoryName      QueryURL = "categoryName"
	Source            QueryURL = "sources"
	Score             QueryURL = "relevanceScore"
	Query             QueryURL = "query"
	Page              QueryURL = "page"
	Limit             QueryURL = "limit"
)

type QueryURL string
