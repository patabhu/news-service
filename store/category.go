package store

import (
	"log"
	"rest-dummy/model"
)

func (d *db) RetrieveByCategory(categoryName string, limit, offset int) ([]*model.NewsArticle, error) {
	var err error
	defer func() {
		if err != nil {
			log.Print(err)
		}
	}()

	news := []*model.NewsArticle{}
	rows, err := d.conn.Query(`
			SELECT 
				id, title, description, llm_summary, url, publication_date, source_name, category, relevance_score, latitude, longitude 
			FROM 
				news_article 
			WHERE 
				$1 = ANY(category)
			ORDER BY 
				publication_date DESC
			LIMIT $2 OFFSET $3`,
		categoryName, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		n := model.NewsArticle{}
		if err = rows.Scan(&n.ID, &n.Title, &n.Description, &n.LLMSummary, &n.URL, &n.PublicationDate, &n.SourceName, &n.Category, &n.RelevanceScore, &n.Latitude, &n.Longitude); err != nil {
			return nil, err
		}
		news = append(news, &n)
	}

	return news, nil
}
