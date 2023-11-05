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
		c.Locals("token", authorization_token)
		return c.Next()
	})
	router_auth.Post("/spotify", func(c *fiber.Ctx) error {
		data := functions.AuthorizeSpotify(c.Query("code"))
		var authformat functions.AuthorizationResponse
		if err := json.Unmarshal([]byte(data), &authformat); err != nil {
			fmt.Println("error unmarshal authorization response")
			fmt.Println(err)
		}
		if len(authformat.Error) > 0 {
			return c.SendString(authformat.Error)
		}
		access_token := authformat.AccessToken
		user := functions.FetchSpotifyUser(access_token)
		var userformat functions.UserResponse
		if err := json.Unmarshal([]byte("{\"user\":"+user+"}"), &userformat); err != nil {
			fmt.Println("error unmarshal usermeta response")
			fmt.Println(err)
		}
		if len(userformat.User.Error) > 0 {
			return c.SendString(userformat.User.Error)
		}
		c.JSON(userformat)
		document, err := functions.FindUserDocumentByID(c, userformat.User.ID)
		if err != nil {
			fmt.Println("error find user document by id")
			fmt.Println(err)
		}
		if err := functions.CreateAuthDocument(document.ID, 1, access_token); err != nil {
			fmt.Println("error create auth document")
			fmt.Println(err)
		}
		database.Redis.Set("key_spotify:"+document.ID, functions.Cryptit(access_token, false))
		database.Redis.Expire("key_spotify:"+document.ID, 3600*time.Second)
		return nil
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
		var docId string
		if _id, err := database.Redis.Get("spotify:" + param); err != nil {
			doc, _ := functions.FindUserDocumentByID(c, param)
			docId = doc.ID
		} else {
			docId = _id
		}
		var token string
		if key, err := database.Redis.Get("key_spotify:" + docId); err != nil {
			fmt.Println("Refreshing token...")
			auth, _ := functions.FindAuthDocumentByRefID(docId, 1)
			token = functions.RefreshToken(auth.Context)
		} else {
			token = key
		}
		c.Locals("token", token)
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
