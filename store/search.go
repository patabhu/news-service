package store

import (
	"log"
	"rest-dummy/model"
)

func (d *db) RetrieveBySearch(query string, limit int) ([]*model.NewsArticle, error) {
	var err error
	defer func() {
		if err != nil {
			log.Print(err)
		}
	}()

	news := []*model.NewsArticle{}
	rows, err := d.conn.Query(`
			WITH ranked_articles AS (
    			SELECT
					id, title, description, llm_summary, url, publication_date, source_name, category, relevance_score, latitude, longitude,
    			    ts_rank_cd(
    			        search_vector,
    			        to_tsquery('english', $1)
    			    ) AS text_score
    			FROM news_article
    			WHERE search_vector @@ to_tsquery('english', $1)
			)
				SELECT
					id, title, description, llm_summary, url, publication_date, source_name, category, relevance_score, latitude, longitude,
				    (
				        text_score * 0.7 +
				        relevance_score * 0.3
				    ) AS final_score
				FROM ranked_articles
				ORDER BY final_score DESC
				LIMIT $2;
			`, query, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		n := model.NewsArticle{}
		var finalScore float64
		if err = rows.Scan(&n.ID, &n.Title, &n.Description, &n.LLMSummary, &n.URL, &n.PublicationDate,
			&n.SourceName, &n.Category, &n.RelevanceScore, &n.Latitude, &n.Longitude, &finalScore); err != nil {
			return nil, err
		}
		news = append(news, &n)
	}

	return news, nil
}
