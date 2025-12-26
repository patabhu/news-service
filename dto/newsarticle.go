package dto

import "time"

type ArticleResponse struct {
	Message  string         `json:"message"`
	Articles []*NewsArticle `json:"articles,omitempty"`
	MetaData MetaData       `json:"meta_data,omitempty"`
}

type NewsArticle struct {
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	URL             string     `json:"url"`
	PublicationDate *time.Time `json:"publication_date"`
	SourceName      string     `json:"source_name"`
	Category        []string   `json:"category"`
	RelevanceScore  float64    `json:"relevance_score"`
	LLMSummary      string     `json:"llm_summary"`
}

type MetaData struct {
	Query          string  `json:"query,omitempty"`
	CategoryName   string  `json:"category_name,omitempty"`
	RelevanceScore float64 `json:"relevanceScore,omitempty"`
	Sources        string  `json:"sources,omitempty"`
	Page           int     `json:"page,omitempty"`
	Limit          int     `json:"limit,omitempty"`
	TotalResults   int     `json:"total_results"`
}
