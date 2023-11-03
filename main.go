package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/tantoony/spotify-status-api-golang/config"
	"github.com/tantoony/spotify-status-api-golang/database"
	"github.com/tantoony/spotify-status-api-golang/functions"

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
		}
		authorization_token := secretKey[0]
		fmt.Println("token: " + authorization_token)
		return c.Next()
	})
	router_auth.Post("/spotify", func(c *fiber.Ctx) error {
		data := functions.AuthorizeSpotify(c.Query("code"))
		var authformat functions.AuthorizationResponse
		err := json.Unmarshal([]byte(data), &authformat)
		if err != nil {
			fmt.Println("error unmarshal authorization response")
			fmt.Println(err)
		}
		access_token := authformat.AccessToken

		user := functions.FetchSpotifyUser(access_token)
		var userformat functions.UserResponse
		err = json.Unmarshal([]byte("{\"user\":"+user+"}"), &userformat)
		if err != nil {
			fmt.Println("error unmarshal usermeta response")
			fmt.Println(err)
		}
		return c.JSON(userformat)
	})

	router_spotify := app.Group("/spotify")
	router_spotify.Use(func(c *fiber.Ctx) error {
		fmt.Println("Spotify middleware")
		return c.Next()
	})

	router_spotify.Get("/:id", func(c *fiber.Ctx) error {
		var j interface{}
		asd := functions.FetchSpotifyUser("BQANOM9e0-nuLkRmClMoZaqvE6KP8HRLjZcNov9fOEfIaAseOPhv5U4F1TQ89sxU16tN-cBWxbarlCd6n2Xh4fnexz_nRJZR4az8rG4_J5zYPhzu4bcTzJuflizG2vqYvuwjq0CD1OgG2TcsfLClxUDkwZmC4ZBI6V1G9rSCKkda9JUw1wwUPHqVLYoAba46DU1D1z3HzHO_tinDFp1cnX7Ug6I")
		err := json.Unmarshal([]byte(asd), &j)
		if err != nil {
			fmt.Println(err)
		}
		return c.JSON(j)
	})

	if err := app.Listen(":5000"); err != nil {
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
