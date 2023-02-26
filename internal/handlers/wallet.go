package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletAddress := vars["address"]
	walletInfo, err := h.services.Wallet.GetByAddress(walletAddress)
	if err != nil {
		errorMessage := fmt.Sprintf("кошелек с адресом %s не найден", walletAddress)
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
