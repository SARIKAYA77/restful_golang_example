package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	."../models"
	"strconv"
	"time"
	"github.com/gorilla/mux"
)
var coinstore = make(map[string]Coin)
var id int = 0 

func Run(){
	http.HandleFunc("/",Handler)
	http.ListenAndServe(":8080",nil)
}

func add_currency(w http.ResponseWriter, r *http.Request) {
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

}

func delete_currencies(w http.ResponseWriter, r *http.Request) {
	vars :=mux.Vars(r)
	key:=vars["id"]
	if _, ok := coinstore[key]; ok {
		delete(coinstore,key)
	}else{
		log.Println("coin not found")
	}
	w.WriteHeader(http.StatusOK)
}	
	

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