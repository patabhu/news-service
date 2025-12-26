package internal

import (
	"rest-dummy/model"
	"rest-dummy/store"
)

func GetNewsByScore(score float64, limit int, offset int) ([]*model.NewsArticle, error) {

	news, err := store.GetNewsReader().RetrieveByScore(float64(score), limit, offset)
	if err != nil {
		return nil, err
	}

	return news, nil
}
