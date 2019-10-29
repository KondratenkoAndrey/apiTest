package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dialect := os.Getenv("db_type")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Создать строку подключения

	connection, err := gorm.Open(dialect, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = connection
	db.Debug().AutoMigrate(&Person{})
}

func GetDb() *gorm.DB {
	return db
}