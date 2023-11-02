package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisInstance struct {
	Client *redis.Client
	ctx    context.Context
}

// Source https://go.dev/tour/methods/1
func (db RedisInstance) get(key string) (string, error) {
	val, err := db.Client.Get(db.ctx, key).Result()
	return val, err
}
func (db RedisInstance) del(keys ...string) error {
	err := db.Client.Del(db.ctx, keys...).Err()
	return err
}
func (db RedisInstance) set(key string, value interface{}) error {
	err := db.Client.Set(db.ctx, key, value, 0).Err()
	return err
}

var (
	Mongo *mongo.Database
	Redis RedisInstance
	app   *fiber.App
)

func main() {
	//fn.GetOs()
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
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	} else {
		fmt.Println("Environment variables successfully loaded...")
	}
	if err := MongoConnection(os.Getenv("MONGO_URI"), os.Getenv("MONGO_DBNAME")); err != nil {
		log.Fatalf("Error connecting to MongoDB")
	} else {
		fmt.Println("MongoDB successfully connected...")
	}
	if err := RedisConnection(os.Getenv("REDIS_URI")); err != nil {
		log.Fatalf("Error connecting to Redis")
	} else {
		fmt.Println("Redis successfully connected...")
	}
	app = fiber.New()
	app.Use(cors.New())
}

func MongoConnection(mongoURI string, dbName string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	Mongo = client.Database(dbName)
	return nil
}

// Source: https://redis.io/docs/clients/go/
func RedisConnection(redisURI string) error {
	opt, err := redis.ParseURL(redisURI)

	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	ctx := context.Background()

	Redis = RedisInstance{
		Client: client,
		ctx:    ctx,
	}
	return nil
}
