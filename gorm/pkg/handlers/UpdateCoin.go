package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/models"
)

func (h handler) UpdateCoin(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedCoin models.Coin
	json.Unmarshal(body, &updatedCoin)

	var coin models.Coin

	if result := h.DB.First(&coin, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	coin.Code = updatedCoin.Code
	coin.Amount = updatedCoin.Amount
	coin.Price = updatedCoin.Price

	h.DB.Save(&coin)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")
}
