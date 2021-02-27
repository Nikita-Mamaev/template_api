package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=nikitamamaev dbname=test sslmode=disable")
	if err != nil {
		panic("DataBase error")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
