package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/db"
	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/handlers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/coin",add_currency).Methods("PUT")
	router.HandleFunc("/coin/{id}",edit_currencies).Methods("PATCH")
	router.HandleFunc("/coin/{id}",delete_currencies).Methods("DELETE")
	router.HandleFunc("/coin/{id}",get_currencies).Methods("GET")
	router.HandleFunc("/coin",get_all_currencies).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080",router))
	log.Println("API is running!")
}