package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/models"
)

func (h handler) GetAllCoin(w http.ResponseWriter, r *http.Request) {
	var coin []models.Coin

	if result := h.DB.Find(&coin); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(coin)
}
