package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)

type Coin struct {
	ID   string `json:id`
	Code string `json:"code"`
	Amount float64 `json:"amount"`
	Current_price float64 `json:"current_price"`
	Old_price float64 `json:"old_price"`
}
var coin []Coin
var coin2 []Coin
func main() {
	paprikaClient := coinpaprika.NewClient(nil)
	var Amount float64
	Amount = 5.4
	var code string
	code = "btc-bitcoin"

	opts := &coinpaprika.PriceConverterOptions{
		BaseCurrencyID: "btc-bitcoin", QuoteCurrencyID: "usd-us-dollars", Amount: Amount,
	}
	result, err := paprikaClient.PriceConverter.PriceConverter(opts)
	if err != nil {
		panic(err)
	}
	//str:=fmt.Sprint(Amount)
	fmt.Printf("str BTC is worth %v US Dollars\n", *result.Price)

	coin = append(coin,
			Coin{ID: "1", Code:code, Amount:Amount, Old_price: *result.Price, Current_price:*result.Price})
	coin2 = append(coin2,
			Coin{ID: "1", Code:code, Amount:Amount, Current_price:*result.Price})

	router := mux.NewRouter()
	router.HandleFunc("/coin",add_currency).Methods("PUT")
	router.HandleFunc("/coin/{id}",edit_currencies).Methods("PATCH")
	router.HandleFunc("/coin/{id}",delete_currencies).Methods("DELETE")
	router.HandleFunc("/coin/{id}",get_currencies).Methods("GET")
	router.HandleFunc("/coin",get_all_currencies).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",router))
}

func add_currency(w http.ResponseWriter, r *http.Request) {
	log.Println("add coin is called")
	var coin Coin
	_ = json.NewDecoder(r.Body).Decode(&coin)

	coin2 = append(coin2, coin)
	json.NewEncoder(w).Encode(coin2)
	
	
}
func edit_currencies(w http.ResponseWriter, r *http.Request) {
	log.Println("edit_currencies is called")
}
func delete_currencies(w http.ResponseWriter, r *http.Request) {
	log.Println("delete_currencies is called")
	params := mux.Vars(r)
	var coin1 Coin

	coin1.ID = params["id"]

	for i, item := range coin2 {
		if item.ID == coin1.ID {
			coin2 = append(coin2[:i], coin2[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(coin2)
}

func get_currencies(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(coin)
	
}
func get_all_currencies(w http.ResponseWriter, r *http.Request) {
	log.Println("get_all_currencies is called")
	json.NewEncoder(w).Encode(coin)
	
}