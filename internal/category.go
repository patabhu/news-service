package internal

import (
	"rest-dummy/model"
	"rest-dummy/store"
)

func GetNewsByCategory(categoryName string, limit, offset int) ([]*model.NewsArticle, error) {

	news, err := store.GetNewsReader().RetrieveByCategory(categoryName, limit, offset)
	if err != nil {
		return nil, err
	}

	return news, nil
}
