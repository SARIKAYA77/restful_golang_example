package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/restful_golang_example/models"
)

func add_currency(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("add coin is called")
	var coin Coin
	err := json.NewDecoder(r.Body).Decode(&coin)
	CheckError(err)
	coin.CreatedOn = time.Now()
	id++
	coin.ID = id
	key := strconv.Itoa(id)
	coinstore[key] = coin
	
	data, err = json.Marshal(coin)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.Header(http.StatusCreated)
	w.Write(data)
}