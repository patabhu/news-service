package internal

import (
	"rest-dummy/model"
	"rest-dummy/store"
)

func GetNewsBySearch(query string, limit int) ([]*model.NewsArticle, error) {

	news, err := store.GetNewsReader().RetrieveBySearch(query, limit)
	if err != nil {
		return nil, err
	}

	return news, nil
}
