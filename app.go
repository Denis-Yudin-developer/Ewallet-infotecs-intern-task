package main

import (
	"Ewallet-infotecs/internal/handlers"
	"Ewallet-infotecs/internal/repository"
	"Ewallet-infotecs/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}
	db, _ := repository.
		NewSqliteDb(viper.GetString("db.driverName"),
			viper.GetString("db.dataSourceName"))

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handlers.NewHandler(services)

	start(handler.InitRoutes())
}

func start(r *mux.Router) {
	logrus.Printf("start application")

	serverAddress := fmt.Sprintf("%s:%s",
		viper.GetString("address"),
		viper.GetString("port"))

	srv := &http.Server{
		Handler:      r,
		Addr:         serverAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logrus.Printf("server is listening port %s", serverAddress)
	logrus.Fatal(srv.ListenAndServe())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
