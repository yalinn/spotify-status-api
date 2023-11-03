package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	CLIENT_ID     string
	CLIENT_SECRET string
	MONGO_URI     string
	MONGO_DBNAME  string
	REDIS_URI     string
)

func InitializeEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	MONGO_URI = os.Getenv("MONGO_URI")
	MONGO_DBNAME = os.Getenv("MONGO_DBNAME")
	REDIS_URI = os.Getenv("REDIS_URI")
}