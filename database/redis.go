package database

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisInstance struct {
	Client *redis.Client
	ctx    context.Context
}

var Redis RedisInstance

// Source https://go.dev/tour/methods/1
func (db RedisInstance) Get(key string) (string, error) {
	val, err := db.Client.Get(db.ctx, key).Result()
	return val, err
}
func (db RedisInstance) Del(keys ...string) error {
	err := db.Client.Del(db.ctx, keys...).Err()
	return err
}
func (db RedisInstance) Set(key string, value interface{}) error {
	err := db.Client.Set(db.ctx, key, value, 0).Err()
	return err
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
