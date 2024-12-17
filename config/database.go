package config

import (
	"fmt"
	"log"
	"product_api/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "rafi:password@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectDB() error {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = Database.AutoMigrate(&entities.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Database connected successfully")

	return nil
}
