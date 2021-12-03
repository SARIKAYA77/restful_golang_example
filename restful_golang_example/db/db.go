package db

import (
	"database/sql"
	"fmt"
	"github.com/restful_golang_example/models"
	"log"
	_ "github.com/lib/pq"

)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "12345"
	dbname = "coindb"
)
var db *sql.DB

func init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password,dbname)
	db, err = sql.Open("postgres", connStr)
	db.SetMaxIdleConns(5)
	if err != nil {
		log.Fatal(err)
	}


}


func NewCoin(data Coin){
	result,err := db.Exec("INSERT INTO coin(code,amount,price) VALUES($1,$2,$3)",data.code,data.amount,data.price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected,err := result.rowsAffected()
	fmt.Printf("etkilenen kayıt sayisi (%d)",rowsAffected)
}

func UpdateCoin(data Coin){
	result,err := db.Exec("UPDATE coin SET code=$2,amount=$3,price=$4 WHERE id=$1",data.ID,data.code,data.amount,data.price)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCoins(data Coin){
	result,err := db.Query("SELECT * FROM coin")
	if err != nil {
		if err = sql.ErrNoRows{
			fmt.Prinln("no record found")
			return
		}
		log.Fatal(err)
	}
	def rows.Close()
	var coin_array []*Coins
	for rows.Next(){
		cn:=$Coin{}
		err :=rows.Scan($cn.ID,$cn.Code,cn.Amount,cn.Price)
		if err != nil {
			log.Fatal(err)
		}
		coin_array = append(coin_array,cn)
	}
	if err = rows.Err(); err!= nil {
		log.Fatal(err)
	}
	for_,cr := range coin_array{
		fmt.Printf("%d - %s, %d, $%.2f",cr.ID,cr.Code,cr.Amount,cn.Price)
	}
}

func GetCoinsbyID(id int){
	var coin string
	err := db.Query("SELECT price FROM coin Where id=$1",id)Scan($coin)
	switch{
	case err==sql.ErrNoRows:
		log.Fatal(err)
	}
	case err!= nil {
		log.Fatal(err)
	default:
		fmt.Printf("coin is",coin)
	}
}