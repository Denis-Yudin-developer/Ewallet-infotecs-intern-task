package handlers

import (
	"Ewallet-infotecs/internal/model"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /send")
	var transaction model.Transaction
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&transaction); err != nil {
		respondWithError(w, http.StatusBadRequest, "Некорректный запрос")
	}
	defer r.Body.Close()

	id, err := h.services.TransactionHistory.Create(transaction)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := map[string]int64{"id": id}
	logrus.Printf("id транзакции: %d", res["id"])
	respondWithJSON(w, http.StatusOK, res)
}

func (h *Handler) GetLast(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /transactions?count=N")
	queryParams := r.URL.Query()
	n, err := strconv.Atoi(queryParams.Get("count"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	transactions, err := h.services.TransactionHistory.LastNTransaction(n)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, transactions)
}
