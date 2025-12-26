package main

import (
	"os"
	"rest-dummy/config"
	"rest-dummy/migration"
)

func main() {
	os.Setenv("GOOGLE_CLOUD_PROJECT", "gen-lang-client-0894469466")
	os.Setenv("GOOGLE_CLOUD_LOCATION", "global")
	os.Setenv("GOOGLE_GENAI_USE_VERTEXAI", "True")

	migration.Migrate(config.Config{
		Postgres: config.PostgresConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "psql",
			DBName:   "testdb",
		},
		Migration: config.MigrationConfig{
			CreateTables: true,
			// APIData:      true,
			APIUrl:       "https://inshorts.com/api/in/en/news?category=top_stories&max_limit=200&include_card_data=true",
			FileLocation: "./migration/newsdump.json",
			FileData:     true,
		},
	})
}
