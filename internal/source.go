package internal

import (
	"rest-dummy/model"
	"rest-dummy/store"
)

func GetNewsBySource(source string, limit, offset int) ([]*model.NewsArticle, error) {

	news, err := store.GetNewsReader().RetrieveBySource(source, limit, offset)
	if err != nil {
		return nil, err
	}

	return news, nil
}
