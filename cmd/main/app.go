package main

import (
	"Ewallet-infotecs/internal/handlers"
	"Ewallet-infotecs/internal/repository"
	"Ewallet-infotecs/internal/service"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

// TODO доставать их из конфига?
const driverName string = "sqlite3"
const dataSourceName string = "../../db.sqlite"

func main() {
	// TODO Добавить конфиги (переменные окружения)
	// TODO добавить логгирование (и пинг дб)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	log.Printf("register wallet handler")
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	start(handler.InitRoutes())
}

func start(r *mux.Router) {
	log.Printf("start application")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("server is listening port 127.0.0.1:8000")
	log.Fatal(srv.ListenAndServe())
}
