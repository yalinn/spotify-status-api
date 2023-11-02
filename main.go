package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
	app := fiber.New()
	app.Use(cors.New())

	if err := app.Listen(":3000"); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
