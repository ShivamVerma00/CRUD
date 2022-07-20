package database

import (
	"CRUD/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connector *gorm.DB

func connect(connectionString string) error {
	var err error

	connector, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		return err
	}
	log.Printf("Connection Successful...")
	connector.AutoMigrate(&model.Contact{})
	return nil
}

func main() {
	config := Config{
		ServerName: "localhost:3306",
		User:       "root",
		Password:   "TFT@us@123",
		DB:         "contacts",
	}

	connectionString := getConnectionString(config)
	err := connect(connectionString)
	if err != nil {
		log.Panic(err.Error())
	}
}
