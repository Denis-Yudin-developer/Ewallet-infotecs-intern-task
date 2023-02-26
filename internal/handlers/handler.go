package handlers

import (
	"Ewallet-infotecs/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/send", h.Send).Methods("POST")
	r.HandleFunc("/api/transactions", h.GetLast).Methods("GET").Queries("count", "{N}")
	r.HandleFunc("/api/wallet/{address}/balance", h.GetBalance).Methods("GET")

	return r
}
