package models

type Coin struct {
	ID   string `json:id`
	Code string `json:"code"`
	Amount float64 `json:"amount"`
	Price float64 `json:"price"`
	Current_price float64 `json:"current_price"`
	Old_price float64 `json:"old_price"`
}