package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/gorilla/mux"
	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/models"
)


var coinstore = make(map[string]Coin)
var id int = 0 

func edit_currencies(w http.ResponseWriter, r *http.Request) {
	log.Println("edit_currencies is called")
	var err error 
	vars:= mux.Vars(r)
	keys := vars["id"]
	var coin_update Coin
	err = json.NewDecoder(r.body).Decode(coin_update)
	CheckError(err)

	if _, ok := coinstore[key]; ok{
		coin_update.ID,_ = strconv.Atoi(key)
		coin_update.ChangedOn = time.Now()
		delete(coinstore,key)
		coinstore[key]=coin_update
	else{
		log.Println("coin not found")
	}
	w.WriteHeader(http.StatusOK)

	}