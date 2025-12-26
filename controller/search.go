package controller

import (
	"net/http"
	"rest-dummy/dto"
	"rest-dummy/internal"
	"strings"
)

func GetNewsBySearch(w http.ResponseWriter, r *http.Request) {
	resp := dto.ArticleResponse{}
	code := http.StatusBadRequest

	defer func() {
		WriteResponse(w, code, resp)
	}()

	query := r.URL.Query().Get(string(Query))
	if query == "" {
		code = http.StatusBadRequest
		resp.Message = "query is required"
		return
	}
	query = strings.ReplaceAll(strings.TrimSpace(strings.ToLower(query)), " ", " | ")

	news, err := internal.GetNewsBySearch(query, NEWS_SEARCH_LIMIT)
	if err != nil {
		code = http.StatusInternalServerError
		resp.Message = "something went wrong"
		return
	}

	data := []*dto.NewsArticle{}

	for _, v := range news {
		data = append(data, &dto.NewsArticle{
			Title:           v.Title,
			Description:     v.Description,
			URL:             v.URL,
			PublicationDate: v.PublicationDate,
			SourceName:      v.SourceName,
			Category:        v.Category,
			RelevanceScore:  v.RelevanceScore,
			LLMSummary:      v.LLMSummary,
		})
	}

	resp.Articles = data
	resp.MetaData.Query = query
	resp.MetaData.TotalResults = len(news)
	code = http.StatusOK
	resp.Message = "success"
}
