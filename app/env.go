package app

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv(filename string) {
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func NewEnv() {
	env := os.Getenv("APP_ENV")
	switch env {
	case "production":
		loadEnv(".env")
		break
	case "development":
		loadEnv(".env.dev")
		break
	default:
		loadEnv(".env.dev")
		break
	}
}
