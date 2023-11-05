package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tantoony/spotify-status-api-golang/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MigrateAuthorizationDocuments(fromRefID string, toRefID string) error {
	from_id, _ := primitive.ObjectIDFromHex(fromRefID)
	to_id, _ := primitive.ObjectIDFromHex(toRefID)
	query := bson.D{{Key: "_id", Value: from_id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "ref", Value: to_id},
		}},
	}
	result, err := database.Auths.UpdateMany(context.TODO(), query, update)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("The number of modified documents in auths collection: %d\n", result.ModifiedCount)
	return nil
}

func RemoveAuthorizationDocuments(refID string, platfom int) error {
	ref_id, _ := primitive.ObjectIDFromHex(refID)
	query := bson.D{
		{Key: "ref", Value: ref_id},
		{Key: "platform", Value: platfom},
	}
	result, err := database.Auths.DeleteMany(context.TODO(), query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("The number of deleted documents in auths collection: %d\n", result.DeletedCount)
	return nil
}

func CreateAuthDocument(refID string, platform int, token string) error {
	ref_id, _ := primitive.ObjectIDFromHex(refID)
	crypted := Cryptit(token, false)
	if _, err := database.Auths.DeleteMany(context.Background(), bson.D{
		{Key: "ref", Value: ref_id},
		{Key: "platform", Value: platform},
	}); err != nil {
		fmt.Println("error delete auth document by ref and platform")
		fmt.Println(err)
		return err
	}
	if _, err := database.Auths.InsertOne(context.TODO(), bson.D{
		{Key: "ref", Value: ref_id},
		{Key: "platform", Value: platform},
		{Key: "token", Value: crypted},
	}); err != nil {
		fmt.Println(err)
		return err
	}
	database.Redis.Set("key_spotify:"+refID, crypted)
	database.Redis.Expire("key_spotify:"+refID, 3600*time.Second)
	return nil
}

func FindUserDocumentByID(c *fiber.Ctx, ID string) (database.UserDocument, error) {
	query := bson.D{{Key: "users", Value: bson.D{{
		Key: "$elemMatch",
		Value: bson.D{
			{Key: "user_id", Value: ID},
			{Key: "platform", Value: 1},
		},
	}}}}
	var result database.UserDocument
	err := database.Profiles.FindOne(context.Background(), query).Decode(&result)
	if err != nil {
		document := new(database.UserDocument)
		document.ID = ""
		document.Users = append([]database.User{}, database.User{
			User_ID:  ID,
			Platform: 1,
		})
		document.AccessKey = c.Locals("token").(string)
		document.CreatedAt = time.Now()
		document.UpdatedAt = time.Now()
		insertionResult, err := database.Profiles.InsertOne(c.Context(), document)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("created new document with id: ", insertionResult.InsertedID)
		result = *document
	}
	database.Redis.Set("spotify:"+ID, result.ID)
	return result, err
}
