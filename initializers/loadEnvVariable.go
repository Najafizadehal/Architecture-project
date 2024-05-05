package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v. Using default values.", err)
	}
}
