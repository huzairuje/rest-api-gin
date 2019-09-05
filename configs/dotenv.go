package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func DotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
