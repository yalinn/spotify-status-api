package main

import (
	"encoding/json"
	"fmt"
	"time"

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
		if len(secretKey) == 0 {
			return c.SendString("No secret key provided")
		}
		if secretKey[0] == "" {
			return c.SendString("Invalid secret key")
		}
		authorization_token := secretKey[0]
		c.Locals("authorization_token", authorization_token)
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
		secretKey := c.GetReqHeaders()["Authorization"]
		if len(secretKey) == 0 {
			return c.SendString("No secret key provided")
		}
		if secretKey[0] == "" {
			return c.SendString("Invalid secret key")
		}
		authorization_token := secretKey[0]
		c.Locals("token", authorization_token)
		return c.Next()
	})
	router_spotify.Use("/:id", func(c *fiber.Ctx) error {
		c.Locals("date", time.Now().String())
		param := c.Params("id")
		fmt.Println(param)
		docId, _ := database.Redis.Get("spotify:" + param)
		fmt.Println(docId)
		return c.Next()
	})

	router_spotify.Get("/", func(c *fiber.Ctx) error {
		s := c.Locals("token").(string)
		return c.SendString(s)
	})

	router_spotify.Get("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.SendString(id)
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
