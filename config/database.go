package config

import (
	"fmt"
	"log"
	"os"
	"product_api/entities"
	"product_api/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() error {
	var err error
	helpers.LoadEnv()
	var DATABASE_URI string = os.Getenv("DATABASE_URI")
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
