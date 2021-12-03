package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/models"
)

func (h handler) DeleteCoin(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])


	var coin models.Coin
	
	if result := h.DB.First(&coin, id); result.Error != nil {
		fmt.Println(result.Error)
	}


	h.DB.Delete(&coin)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")
}
