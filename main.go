package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/tantoony/spotify-status-api/config"
	"github.com/tantoony/spotify-status-api/database"
	"github.com/tantoony/spotify-status-api/functions"

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
		var docId string
		if _id, err := database.Redis.Get("spotify:" + param); err != nil {
			doc, _ := functions.FindUserDocumentByID(c, param)
			docId = doc.ID
		} else {
			docId = _id
		}
		var token string
		if key, err := database.Redis.Get("key_spotify:" + docId); err != nil || len(key) == 0 {
			fmt.Println("Refreshing token...")
			auth, _ := functions.FindAuthDocumentByRefID(docId, 1)
			token = functions.RefreshToken(auth.Context)
		} else {
			token = functions.Cryptit(key, true)
		}
		c.Locals("token", token)
		return c.Next()
	})

	router_spotify.Get("/", func(c *fiber.Ctx) error {
		s := c.Locals("token").(string)
		return c.SendString(s)
	})

	router_spotify.Get("/:id", func(c *fiber.Ctx) error {
		//id := c.Params("id")
		now_playing_response := functions.UserPlaying(c.Locals("token").(string))
		if len(now_playing_response) == 0 {
			fmt.Println("asd")
			return c.SendString("No song playing")
		}
		var now_playing functions.UserPlayingResponse
		if err := json.Unmarshal([]byte(now_playing_response), &now_playing); err != nil {
			fmt.Println("error unmarshal now_playing response")
			fmt.Println(err)
		}
		queue_response := functions.UserQueue(c.Locals("token").(string))
		var queue functions.UserPlayerQueueResponse
		if len(queue_response) == 0 {
			return c.SendString("No song in queue")
		}
		if err := json.Unmarshal([]byte(queue_response), &queue); err != nil {
			fmt.Println("error unmarshal queue response")
			fmt.Println(err)
		}
		var response = new(functions.SpotifyResponse)
		response.IsActive = now_playing.IsPlaying
		response.Type = now_playing.Device.Type
		response.ShuffleState = now_playing.ShuffleState
		response.RepeatState = now_playing.RepeatState
		response.IsPlaying = now_playing.IsPlaying
		response.TimeStamp = time.Now().UnixMilli()
		response.Song = now_playing.Item.Name
		response.Progress.From = GTS(now_playing.ProgressMs)
		response.Progress.To = GTS(now_playing.Item.DurationMs)
		response.Artists = make([]functions.Artist, 0)
		for _, artist := range now_playing.Item.Artists {
			response.Artists = append(response.Artists, functions.Artist{
				Name: artist.Name,
				Url:  artist.ExternalURLs.Spotify,
			})
		}
		response.ProgressMs = now_playing.ProgressMs
		response.DurationMs = now_playing.Item.DurationMs
		response.Image.Url = now_playing.Item.Album.Images[0].URL
		response.Image.Height = now_playing.Item.Album.Images[0].Height
		response.Image.Width = now_playing.Item.Album.Images[0].Width
		response.Url = now_playing.Item.ExternalURLs.Spotify
		response.ReqTime = c.Locals("date").(string)
		response.Queue = make([]functions.QueuedSong, 0)
		for _, item := range queue.Queue {
			response.Queue = append(response.Queue, functions.QueuedSong{
				Name:    item.Name,
				Artists: item.Artists[0].Name,
				Image: functions.Image{
					Url:    item.Album.Images[0].URL,
					Height: item.Album.Images[0].Height,
					Width:  item.Album.Images[0].Width,
				},
			})
		}
		//response.Queue = response.Queue[1:]
		return c.JSON(response)
	})

	if err := app.Listen(config.PORT); err != nil {
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

func GTS(ms int) string {
	seconds := ms / 1000
	minutes := seconds / 60
	seconds = seconds % 60
	var str string
	if seconds < 10 {
		str = "0"
	} else {
		str = ""
	}
	return fmt.Sprintf("%d:%s%d", minutes, str, seconds)
}
