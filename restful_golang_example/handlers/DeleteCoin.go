package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/restful_golang_example/models"
)

var coinstore = make(map[string]Coin)
var id int = 0 

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
	