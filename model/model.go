package model

import "time"

type NewsArticle struct {
	ID              string
	Title           string
	Description     string
	LLMSummary      string
	URL             string
	PublicationDate *time.Time
	SourceName      string
	Category        []string
	RelevanceScore  float64
	Latitude        float64
	Longitude       float64
}
