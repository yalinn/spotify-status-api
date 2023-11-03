package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	} else {
		fmt.Println("Environment variables successfully loaded...")
	}

}
