package controller

import (
	"net/http"
	"rest-dummy/dto"
	"rest-dummy/internal"
	"strconv"
	"strings"
)

func GetNewsByCategory(w http.ResponseWriter, r *http.Request) {
	resp := dto.ArticleResponse{}
	code := http.StatusBadRequest

	defer func() {
		WriteResponse(w, code, resp)
	}()
	urlQuery := r.URL.Query()

	categoryName := urlQuery.Get(string(CategoryName))
	if categoryName == "" {
		code = http.StatusBadRequest
		resp.Message = "categoryName is required"
		return
	}
	categoryName = strings.TrimSpace(strings.ToLower(categoryName))

	page := NEWS_PAGE
	if p, err := strconv.Atoi(urlQuery.Get(string(Page))); err == nil && p > 0 {
		page = p
	}

	limit := NEWS_LIMIT
	if l, err := strconv.Atoi(urlQuery.Get(string(Limit))); err == nil && l < NEWS_LIMIT {
		limit = l
	}
	offset := (page - 1) * limit

	news, err := internal.GetNewsByCategory(categoryName, limit, offset)
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
	resp.MetaData.CategoryName = categoryName
	resp.MetaData.Limit = limit
	resp.MetaData.Page = page
	resp.MetaData.TotalResults = len(data)
	code = http.StatusOK
	resp.Message = "success"
}
