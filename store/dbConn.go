package store

import (
	"rest-dummy/config"

	"github.com/jackc/pgx"
)

var dbconn *db

type db struct {
	conn *pgx.Conn
}

func Init(conf config.Config) {
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
	dbconn = &db{conn}
}
