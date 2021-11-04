package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() (*gorm.DB, error) {
	var db *gorm.DB
	godotenv.Load()

	dbUser, isUser := os.LookupEnv("DB_USER")
	dbPass, isPass := os.LookupEnv("DB_PASSWORD")
	dbHost, isHost := os.LookupEnv("DB_HOST")

	if !isUser || !isPass || !isHost {
		fmt.Println("Cant read .env file")
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=postgres sslmode=disable`, dbHost, dbUser, dbPass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	err = db.AutoMigrate(&User{}, &Post{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
