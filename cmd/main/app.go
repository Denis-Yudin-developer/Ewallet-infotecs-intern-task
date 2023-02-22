package main

import (
	"Ewallet-infotecs/internal/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Printf("create router")
	r := mux.NewRouter()

	log.Printf("register user handler")
	handler := user.NewHandler()
	handler.Register(r)

	start(r)
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
