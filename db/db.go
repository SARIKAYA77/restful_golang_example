package db

import (
	"log"

	"github.com/SARIKAYA77/restful_golang_example/restful_golang_example/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://pg:12345@localhost:5432/restful_golang_example"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}