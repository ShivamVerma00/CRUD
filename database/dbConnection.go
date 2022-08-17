package database

import (
	"CRUD/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB

var dbURL = "root:TFT@us@123@tcp(127.0.0.1:3306)/contact"

func init() {
	var err error

	connector, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Error in DB Connection", err)
	}
	log.Printf("Connection Successful...")
	connector.AutoMigrate(&model.Contact{})
}
