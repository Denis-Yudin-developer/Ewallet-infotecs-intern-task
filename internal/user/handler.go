package user

import (
	"Ewallet-infotecs/internal/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/api/send", h.Send).Methods("POST")
	router.HandleFunc("/api/transactions", h.GetLast).Methods("GET").Queries("count", "{N}")
	router.HandleFunc("/api/wallet/{address}/balance", h.GetBalance).Methods("GET")
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User address is %s", vars["address"])
}

func (h *handler) Send(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post works")
}

func (h *handler) GetLast(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("count")
	fmt.Fprintf(w, "last transactions: %s", key)
}
