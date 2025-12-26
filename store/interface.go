package store

import "rest-dummy/model"

type NewsReader interface {
	RetrieveByCategory(categoryName string, limit, offset int) ([]*model.NewsArticle, error)
	RetrieveByScore(score float64, limit, offset int) ([]*model.NewsArticle, error)
	RetrieveBySource(sourceName string, limit, offset int) ([]*model.NewsArticle, error)
	RetrieveBySearch(query string, limit int) ([]*model.NewsArticle, error)
}

func GetNewsReader() NewsReader {
	return dbconn
}
