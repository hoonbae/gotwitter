package core

import (
	"fmt"
	"log"

	"gotwitter/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Storage StorageStruct

type StorageStruct struct {
	DB *gorm.DB
}

func (s *StorageStruct) Start() {
	initDatabase()
}

// Database

func initDatabase() {
	log.Println("START: DATABASE INITIALIZING")
	var err error

	Storage.DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=gotwitter sslmode=disable")

	if err != nil {
		log.Fatal(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	Storage.DB.LogMode(true)

	Storage.DB.AutoMigrate(
		&models.User{},
		&models.Tweet{},
	)

	log.Println("END: DATABASE INITIALIZING")
}
