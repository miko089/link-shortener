package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetBaseURL() string {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	return os.Getenv("BASE_URL")
}
