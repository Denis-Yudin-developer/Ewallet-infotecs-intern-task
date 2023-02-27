package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	logrus.Print("Запрос на маршрут /wallet/{address}/balance")
	vars := mux.Vars(r)
	walletAddress := vars["address"]
	walletInfo, err := h.services.Wallet.GetByAddress(walletAddress)
	if err != nil {
		errorMessage := fmt.Sprintf("кошелек не найден")
		respondWithError(w, http.StatusNotFound, errorMessage)
		return
	}
	allBytes, err := json.Marshal(walletInfo)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(allBytes)
}
