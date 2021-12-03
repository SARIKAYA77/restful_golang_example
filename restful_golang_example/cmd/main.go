package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restful_golang_example/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/coin",handlers.AddCoin).Methods("PUT")
	router.HandleFunc("/coin/{id}",handlers.UpdateCoin).Methods("PATCH")
	router.HandleFunc("/coin/{id}",handlers.DeleteCoin).Methods("DELETE")
	router.HandleFunc("/coin/{id}",handlers.GetCoin).Methods("GET")
	router.HandleFunc("/coin",handlers.GetAllCoin).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080",router))
	log.Println("API is running!")
}