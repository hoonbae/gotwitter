package core

import (
	"fmt"
	"os"

	"twitter/models"

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

	Storage.DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	Storage.DB.LogMode(Config.Local())

	Storage.DB.AutoMigrate(
		&models.User{},
		&models.Tweet{},
	)

	log.Println("END: DATABASE INITIALIZING")
}
