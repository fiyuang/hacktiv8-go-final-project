package database

import (
	"fmt"
	"hacktiv8-go-final-project/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbport := os.Getenv("DBPORT")
	dbname := os.Getenv("DBNAME")

	// fmt.Println(dbname)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	db.Debug().AutoMigrate(models.User{})
}

func GetDB() *gorm.DB {
	return db
}
