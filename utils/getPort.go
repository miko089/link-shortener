package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}
	if os.Getenv("PORT") == "" {
		return ":80"
	}
	return ":" + os.Getenv("PORT")
}
