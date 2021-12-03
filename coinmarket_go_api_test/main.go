package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"time"
	"os"
	"net/url"
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	
	// "github.com/coinpaprika/coinpaprika-api-go-client/coinpaprika"
)
const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "12345"
	dbname = "coindb"
)

// type Coin struct {
// 	ID   string `json:id`
// 	Code string `json:"code"`
// 	Amount float64 `json:"amount"`
// 	Current_price float64 `json:"current_price"`
// 	Old_price float64 `json:"old_price"`
// }


type coinList struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
		TotalCount   int         `json:"total_count"`
	} `json:"status"`
	Data []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		DateAdded         time.Time   `json:"date_added"`
		Tags              []string    `json:"tags"`
		MaxSupply         int         `json:"max_supply"`
		CirculatingSupply int         `json:"circulating_supply"`
		TotalSupply       int         `json:"total_supply"`
		Platform          interface{} `json:"platform"`
		CmcRank           int         `json:"cmc_rank"`
		LastUpdated       time.Time   `json:"last_updated"`
		Quote             struct {
			Usd struct {
				Price            float64   `json:"price"`
				Volume24H        float64   `json:"volume_24h"`
				PercentChange1H  float64   `json:"percent_change_1h"`
				PercentChange24H float64   `json:"percent_change_24h"`
				PercentChange7D  float64   `json:"percent_change_7d"`
				PercentChange30D float64   `json:"percent_change_30d"`
				PercentChange60D float64   `json:"percent_change_60d"`
				PercentChange90D float64   `json:"percent_change_90d"`
				MarketCap        float64   `json:"market_cap"`
				LastUpdated      time.Time `json:"last_updated"`
			} `json:"USD"`
		} `json:"quote"`
	} `json:"data"`
}

type coinPricing struct {
	Price            float64 `json:"price"`
	Volume24H        float64
	PercentChange1H  float64
	PercentChange24H float64
	PercentChange7D  float64
	PercentChange30D float64
	PercentChange60D float64
	PercentChange90D float64
	MarketCap        float64
	LastUpdated      time.Time
}

type apiExplorer struct {
	Address string `json:"address"`
	Eth     struct {
		Price struct {
			Rate            float64 `json:"rate"`
			Diff            float64 `json:"diff"`
			Diff7D          float64 `json:"diff7d"`
			Ts              int     `json:"ts"`
			MarketCapUsd    float64 `json:"marketCapUsd"`
			AvailableSupply float64 `json:"availableSupply"`
			Volume24H       float64 `json:"volume24h"`
			Diff30D         float64 `json:"diff30d"`
			VolDiff1        float64 `json:"volDiff1"`
			VolDiff7        float64 `json:"volDiff7"`
			VolDiff30       float64 `json:"volDiff30"`
		} `json:"price"`
		RawBalance float64 `json:"rawBalance"`
	} `json:"ETH"`
	CountTxs int `json:"countTxs"`
	Tokens   []struct {
		TokenInfo struct {
			Address           string `json:"address"`
			Name              string `json:"name"`
			Decimals          string `json:"decimals"`
			Symbol            string `json:"symbol"`
			TotalSupply       string `json:"totalSupply"`
			Owner             string `json:"owner"`
			LastUpdated       int    `json:"lastUpdated"`
			IssuancesCount    int    `json:"issuancesCount"`
			HoldersCount      int    `json:"holdersCount"`
			Description       string `json:"description"`
			Website           string `json:"website"`
			Twitter           string `json:"twitter"`
			Reddit            string `json:"reddit"`
			Telegram          string `json:"telegram"`
			Image             string `json:"image"`
			Coingecko         string `json:"coingecko"`
			EthTransfersCount int    `json:"ethTransfersCount"`
			Price             struct {
				Rate            float64 `json:"rate"`
				Diff            float64 `json:"diff"`
				Diff7D          float64 `json:"diff7d"`
				Ts              int     `json:"ts"`
				MarketCapUsd    float64 `json:"marketCapUsd"`
				AvailableSupply int     `json:"availableSupply"`
				Volume24H       float64 `json:"volume24h"`
				Diff30D         float64 `json:"diff30d"`
				VolDiff1        float64 `json:"volDiff1"`
				VolDiff7        float64 `json:"volDiff7"`
				VolDiff30       float64 `json:"volDiff30"`
				Currency        string  `json:"currency"`
			} `json:"price"`
			PublicTags []string `json:"publicTags"`
		} `json:"tokenInfo,omitempty"`
		Balance    float32 `json:"balance"`
		TotalIn    int64   `json:"totalIn"`
		TotalOut   int64   `json:"totalOut"`
		RawBalance string  `json:"rawBalance"`
	} `json:"tokens"`
}

type tokenInfo struct {
	Address           string  `json:"address"`
	Name              string  `json:"name"`
	Decimals          string  `json:"decimals"`
	Symbol            string  `json:"symbol"`
	TotalSupply       string  `json:"totalSupply"`
	Owner             string  `json:"owner"`
	LastUpdated       int     `json:"lastUpdated"`
	IssuancesCount    int     `json:"issuancesCount"`
	HoldersCount      int     `json:"holdersCount"`
	Description       string  `json:"description"`
	Website           string  `json:"website"`
	Twitter           string  `json:"twitter"`
	Reddit            string  `json:"reddit"`
	Telegram          string  `json:"telegram"`
	Image             string  `json:"image"`
	Coingecko         string  `json:"coingecko"`
	EthTransfersCount int     `json:"ethTransfersCount"`
	Balance           float32 `json:"balance"`
	RawBalance        string  `json:"rawbalance"`
	Price             struct {
		Rate            float64 `json:"rate"`
		Diff            float64 `json:"diff"`
		Diff7D          float64 `json:"diff7d"`
		Ts              int     `json:"ts"`
		MarketCapUsd    float64 `json:"marketCapUsd"`
		AvailableSupply int     `json:"availableSupply"`
		Volume24H       float64 `json:"volume24h"`
		Diff30D         float64 `json:"diff30d"`
		VolDiff1        float64 `json:"volDiff1"`
		VolDiff7        float64 `json:"volDiff7"`
		VolDiff30       float64 `json:"volDiff30"`
		Currency        string  `json:"currency"`
	}
}
func main() {
  client := &http.Client{}
  req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
  if err != nil {
    log.Print(err)
    os.Exit(1)
  }

  q := url.Values{}
  q.Add("start", "1")
  q.Add("limit", "10")
  q.Add("convert", "USD")

  req.Header.Set("Accepts", "application/json")
  req.Header.Add("X-CMC_PRO_API_KEY", "ec14dbba-fb18-47e5-8557-d010ffa72525")
  req.URL.RawQuery = q.Encode()

  resp, err := client.Do(req);
  if err != nil {
    fmt.Println("Error sending request to server")
    os.Exit(1)
  }
  fmt.Println(resp.Status);
  respBody, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(respBody));
 
  var errs error
  connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password,dbname)
  db, errs := sql.Open("postgres", connStr)
  db.SetMaxIdleConns(5)
  if err != nil {
	log.Fatal(errs)
  }


  fmt.Println("Successfully created connection to database")
  var apiResponse coinList
  json.Unmarshal(respBody, &apiResponse)
  fmt.Println(apiResponse.Data[0].Quote.Usd[0].Price);
// Insert some data into table.

}

// }
// var coin []Coin
// var coin2 []Coin
// 	paprikaClient := coinpaprika.NewClient(nil)
// 	var Amount float64
// 	Amount = 5.4
// 	var code string
// 	code = "btc-bitcoin"

// 	opts := &coinpaprika.PriceConverterOptions{
// 		BaseCurrencyID: "btc-bitcoin", QuoteCurrencyID: "usd-us-dollars", Amount: Amount,
// 	}
// 	result, err := paprikaClient.PriceConverter.PriceConverter(opts)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//str:=fmt.Sprint(Amount)
// 	fmt.Printf("str BTC is worth %v US Dollars\n", *result.Price)

// 	coin = append(coin,
// 			Coin{ID: "1", Code:code, Amount:Amount, Old_price: *result.Price, Current_price:*result.Price})
// 	coin2 = append(coin2,
// 			Coin{ID: "1", Code:code, Amount:Amount, Current_price:*result.Price})

// 	router := mux.NewRouter()
// 	router.HandleFunc("/coin",add_currency).Methods("PUT")
// 	router.HandleFunc("/coin/{id}",edit_currencies).Methods("PATCH")
// 	router.HandleFunc("/coin/{id}",delete_currencies).Methods("DELETE")
// 	router.HandleFunc("/coin/{id}",get_currencies).Methods("GET")
// 	router.HandleFunc("/coin",get_all_currencies).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8080",router))
// }
// func main() {
// 	response, err := http.Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if response.StatusCode == 200 {
// 		data, _ := ioutil.ReadAll(response.Body)
// 		//fmt.Println(string(data))

// 		var apiResponse ApiResponse
// 		json.Unmarshal(data, &apiResponse)

// 		//fmt.Println(apiResponse)

// 		ac := accounting.Accounting{Symbol: "$", Precision: 4}
// 		fmt.Println("Last Updated:", apiResponse.Time.Updated, " Bitcoin (USD): ", ac.FormatMoney(apiResponse.Bpi.Usd.Rate_Float))

// 	} else {
// 		fmt.Println("Http call failed with status: ", response.Status)
// 	}

// }

// func add_currency(w http.ResponseWriter, r *http.Request) {
// 	log.Println("add coin is called")
// 	var coin Coin
// 	_ = json.NewDecoder(r.Body).Decode(&coin)

// 	coin2 = append(coin2, coin)
// 	json.NewEncoder(w).Encode(coin2)
	
	
// }
// func edit_currencies(w http.ResponseWriter, r *http.Request) {
// 	log.Println("edit_currencies is called")
// }
// func delete_currencies(w http.ResponseWriter, r *http.Request) {
// 	log.Println("delete_currencies is called")
// 	params := mux.Vars(r)
// 	var coin1 Coin

// 	coin1.ID = params["id"]

// 	for i, item := range coin2 {
// 		if item.ID == coin1.ID {
// 			coin2 = append(coin2[:i], coin2[i+1:]...)
// 		}
// 	}

// 	json.NewEncoder(w).Encode(coin2)
// }

// func get_currencies(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(coin)
	
// }
// func get_all_currencies(w http.ResponseWriter, r *http.Request) {
// 	log.Println("get_all_currencies is called")
// 	json.NewEncoder(w).Encode(coin)
	
// }