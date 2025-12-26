package migration

import (
	"rest-dummy/config"

	"github.com/jackc/pgx"
)

func Connect(conf config.Config) *pgx.Conn {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     conf.Postgres.Host,
		Port:     conf.Postgres.Port,
		Database: conf.Postgres.DBName,
		User:     conf.Postgres.User,
		Password: conf.Postgres.Password,
	})
	if err != nil {
		panic(err)
	}
	return conn
}
