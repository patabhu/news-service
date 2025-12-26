package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-dummy/config"
	"rest-dummy/controller"
	"rest-dummy/store"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	psqlPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	store.Init(config.Config{
		Postgres: config.PostgresConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     uint16(psqlPort),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
		},
	})

	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	newsRouteV1 := api.PathPrefix("/v1/news/").Subrouter()

	newsRouteV1.HandleFunc("/category", controller.GetNewsByCategory).Methods(http.MethodGet)
	newsRouteV1.HandleFunc("/source", controller.GetNewsBySource).Methods(http.MethodGet)
	newsRouteV1.HandleFunc("/score", controller.GetNewsByScore).Methods(http.MethodGet)
	newsRouteV1.HandleFunc("/search", controller.GetNewsBySearch).Methods(http.MethodGet)

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	log.Print("server running ", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
