package functions

import (
	"context"
	"fmt"

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
