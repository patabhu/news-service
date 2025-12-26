package main

import (
	"fmt"
	"os"
	"rest-dummy/config"
	"rest-dummy/migration"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(godotenv.Load())

	psqlPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	createTables, _ := strconv.ParseBool(os.Getenv("CREATE_TABLES"))
	apiDataMigration, _ := strconv.ParseBool(os.Getenv("API_DATA_MIGRATION"))
	fileDataMigration, _ := strconv.ParseBool(os.Getenv("FILE_DATA_MIGRATION"))

	migration.Migrate(config.Config{
		Postgres: config.PostgresConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(psqlPort),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
		Migration: config.MigrationConfig{
			CreateTables: createTables,
			APIData:      apiDataMigration,
			APIUrl:       "https://inshorts.com/api/in/en/news?category=top_stories&max_limit=200&include_card_data=true",
			FileLocation: "./migration/newsdump.json",
			FileData:     fileDataMigration,
		},
	})
}
