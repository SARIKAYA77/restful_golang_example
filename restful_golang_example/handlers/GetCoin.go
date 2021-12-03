package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/models"
)
var coinstore = make(map[string]Coin)
var id int = 0 

func get_currencies(w http.ResponseWriter, r *http.Request) {
	log.Println("get_currencies is called")
	var coin Coin
	vars := mux.Vars(r)
	key, _ = strconv.Atoi(vars["id"])
	for _, item = range coinstore{
		if item.ID == key {
			coin = item
		}
	}
	data, err = json.Marshal(coin)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.Header(http.StatusCreated)
	w.Write(data)
	
}