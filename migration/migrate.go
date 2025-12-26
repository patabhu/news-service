package migration

import (
	"rest-dummy/config"
	"rest-dummy/model"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

func Migrate(conf config.Config) {

	migrateConf := conf.Migration
	dbconn := Connect(conf)

	if migrateConf.CreateTables {
		createTables(dbconn)
	}

	if migrateConf.APIData {
		newsDump := getAPINewsDump(migrateConf.APIUrl)
		dumpData(dbconn, newsDump)
	}

	if migrateConf.FileData {
		newsDump := getNewsFileDump(migrateConf.FileLocation)
		dumpData(dbconn, newsDump)
	}

}

func createTables(dbConn *pgx.Conn) {
	_, err := dbConn.Exec(`CREATE table news_article(
								id uuid,
								title varchar(200),
								description text,
								llm_summary text,
								url varchar(500),
								publication_date timestamp,
								source_name	varchar(50),
								category TEXT[],
								relevance_score real,
								latitude real,
								longitude real
				);`)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		panic(err)
	}

	_, err = dbConn.Exec(`ALTER TABLE news_article
				ADD COLUMN search_vector tsvector
				GENERATED ALWAYS AS (
				  to_tsvector('english', coalesce(title,'') || ' ' || coalesce(description,''))
				) STORED;`)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		panic(err)
	}

	_, err = dbConn.Exec(`CREATE INDEX idx_news_search
				ON news_article
				USING GIN (search_vector);`)
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		panic(err)
	}

}

func dumpData(dbConn *pgx.Conn, newsList *newsData) {

	for _, v := range newsList.Data.NewsList {
		newobj := v.NewsObject

		categoryNames := newobj.CategoryNames
		for k, v := range newobj.CategoryNames {
			newobj.CategoryNames[k] = strings.ToLower(v)
		}

		if len(newobj.CategoryNames) == 0 {
			categoryNames = []string{"general"}
		}

		publicationDate := time.UnixMilli(newobj.CreatedAt)

		summary := generateSummary(newobj.Content)

		n := model.NewsArticle{
			ID:              uuid.NewString(),
			Title:           newobj.Title,
			Description:     newobj.Content,
			LLMSummary:      summary,
			URL:             newobj.SourceURL,
			PublicationDate: &publicationDate,
			SourceName:      strings.ToLower(newobj.SourceName),
			Category:        categoryNames,
			RelevanceScore:  newobj.ImpressiveScore,
		}

		_, err := dbConn.Exec(`INSERT INTO news_article(
			id, title, description, llm_summary, url, publication_date, 
			source_name, category, relevance_score, latitude, longitude 
		) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
			n.ID,
			n.Title,
			n.Description,
			n.LLMSummary,
			n.URL,
			n.PublicationDate,
			n.SourceName,
			n.Category,
			n.RelevanceScore,
			n.Latitude,
			n.Longitude)
		if err != nil {
			panic(err)
		}
	}

}
