package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/restful_golang_example/models"
)
var coinstore = make(map[string]Coin)
var id int = 0 

func get_all_currencies(w http.ResponseWriter, r *http.Request) {
	log.Println("get_all_currencies is called")
	var coins []Coin
	for _ , coin = range coinstore{
		coin = append(coins,coin)
	}
	data, err = json.Marshal(coins)
	CheckError(err)
	w.Header().Set("Content-Type", "application/json")
	w.Header(http.StatusOK)
	w.Write(data)
	
}
