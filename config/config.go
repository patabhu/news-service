package config

type Config struct {
	Postgres  PostgresConfig
	Migration MigrationConfig
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	DBName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type MigrationConfig struct {
	CreateTables bool
	FileData     bool
	FileLocation string
	APIData      bool
	APIUrl       string
}
