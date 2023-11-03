package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/tantoony/spotify-status-api-golang/config"
	"github.com/tantoony/spotify-status-api-golang/database"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	app *fiber.App
)

func main() {
	router_auth := app.Group("/auth")
	router_auth.Use(func(c *fiber.Ctx) error {
		secretKey := c.GetReqHeaders()["Authorization"]
		if len(secretKey) == 0 || secretKey[0] == "" {
			return c.SendString("No secret key provided")
		} else {
			secretKey := secretKey[0]
			fmt.Println(secretKey)
		}
		fmt.Println(secretKey)
		return c.Next()
	})
	router_auth.Post("/spotify", func(c *fiber.Ctx) error {
		return c.SendString("[POST]/auth/spotify")
	})
	router_spotify := app.Group("/spotify")
	router_spotify.Use(func(c *fiber.Ctx) error {
		fmt.Println("Spotify middleware")
		return c.Next()
	})

	router_spotify.Get("/:id", func(c *fiber.Ctx) error {
		return c.SendString("[GET]/spotify/:id")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Info("Oops... Server is not running! Reason: %v", err)
	} else {
		log.Info("Server is running on port 3000...")
	}

}

func init() {
	config.InitializeEnv()
	if err := database.MongoConnection(config.MONGO_URI, config.MONGO_DBNAME); err != nil {
		log.Fatalf("Error connecting to MongoDB")
	} else {
		fmt.Println("MongoDB successfully connected...")
	}
	if err := database.RedisConnection(config.REDIS_URI); err != nil {
		log.Fatalf("Error connecting to Redis")
	} else {
		fmt.Println("Redis successfully connected...")
	}
	app = fiber.New()
	app.Use(cors.New())
}
