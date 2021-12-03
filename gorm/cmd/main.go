package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/restful_golang_example/restful_golang_example/db"
	"github.com/restful_golang_example/restful_golang_example/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/coin",h.AddCoin()).Methods("PUT")
	router.HandleFunc("/coin/{id}",h.UpdateCoin(id).Methods("")).Methods("PATCH")
	router.HandleFunc("/coin/{id}",h.DeleteCoin()).Methods("DELETE")
	router.HandleFunc("/coin/{id}",h.GetCoin()).Methods("GET")
	router.HandleFunc("/coin",h.GetAllCoin()).Methods("GET")

	log.Println("API is running!")
	http.ListenAndServe(":8080", router)
}
