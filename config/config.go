package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var ApiKey string

func LoadEnv() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	ApiKey = os.Getenv("ALPHA_VANTAGE_API_KEY")
}
