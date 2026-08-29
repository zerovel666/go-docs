package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DB          string
	DB_USER     string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_PASSWORD string
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB = os.Getenv("DB")
	DB_USER = os.Getenv("DB_USER")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
}
