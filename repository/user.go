package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"user-vote/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
func getCollection() *mongo.Collection {
	return database.DB.Collection("user")
}*/

func CreateUser(user *domain.User, db *mongo.Database) error {
	fmt.Println("CreateUser: ")
	//fmt.Printf("%+v", database.GetCollectionUser())
	collectionUser := db.Collection("user")
	fmt.Printf("%+v", user)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := collectionUser.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func GetAllUsers(db *mongo.Database) ([]*domain.User, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	return filterUsers(filter, db)
}

func filterUsers(filter interface{}, db *mongo.Database) ([]*domain.User, error) {
	// A slice of users for storing the decoded documents
	var users []*domain.User
	var ctx = context.TODO()
	collectionUser := db.Collection("user")
	cur, err := collectionUser.Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var t domain.User
		err := cur.Decode(&t)
		if err != nil {
			return users, err
		}

		users = append(users, &t)
	}

	if err := cur.Err(); err != nil {
		return users, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	return users, nil
}

func UpdateUser(user *domain.User, db *mongo.Database) error {
	filter := bson.M{"_id": bson.M{"$eq": user.ID}}
	var ctx = context.TODO()
	// Create a nested BSON document for the documents' fields that are updated
	update := bson.M{
		"$set": bson.M{
			"name":            user.Name,
			"dateBirth":       false,
			"address.street":  user.Address.Street,
			"address.number":  user.Address.Number,
			"address.county":  user.Address.County,
			"address.city":    user.Address.City,
			"address.state":   user.Address.State,
			"address.zipCode": user.Address.ZipCode,
		},
	}
	collectionUser := db.Collection("user")
	_, err := collectionUser.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

/*
filter := bson.M{"_id": bson.M{"$eq": objID}}
*/

func DeleteUser(id string, db *mongo.Database) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	var ctx = context.TODO()
	collectionUser := db.Collection("user")
	res, err := collectionUser.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No user were deleted")
	}

	return nil
}
