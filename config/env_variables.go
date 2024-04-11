package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SPOTIFY_CLIENT_ID     string
	SPOTIFY_CLIENT_SECRET string
	MONGO_URI             string
	MONGO_DBNAME          string
	REDIS_URI             string
	REDIRECT_URI          string
	PORT                  string = "5871"
)

func InitializeEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	SPOTIFY_CLIENT_ID = os.Getenv("SPOTIFY_CLIENT_ID")
	SPOTIFY_CLIENT_SECRET = os.Getenv("SPOTIFY_CLIENT_SECRET")
	MONGO_URI = os.Getenv("MONGO_URI")
	MONGO_DBNAME = os.Getenv("MONGO_DBNAME")
	REDIS_URI = os.Getenv("REDIS_URI")
	REDIRECT_URI = os.Getenv("REDIRECT_URI")
}
