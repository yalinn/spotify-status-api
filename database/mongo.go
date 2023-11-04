package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Mongo    *mongo.Database
	Profiles *mongo.Collection
	Auths    *mongo.Collection
)

func MongoConnection(mongoURI string, dbName string) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	Mongo = client.Database(dbName)
	Profiles = Mongo.Collection("app.profiles")
	Auths = Mongo.Collection("app.auths")
	return nil
}

type User struct {
	User_ID  string  `json:"user_id"`
	Platform float64 `json:"platform"`
}

type AuthDocument struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Platform  float64   `json:"platform"`
	Context   string    `json:"context"`
	Reference string    `json:"ref"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserDocument struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Users     []User    `json:"users"`
	AccessKey string    `json:"accessKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
