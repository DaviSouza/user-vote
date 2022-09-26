package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDataBase() *mongo.Database {
	uri := "mongodb+srv://admin:GBJ7zdMUxwqfxeC4@cluster0.4yexf9q.mongodb.net/?retryWrites=true&w=majority"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	database := client.Database("user_vote_game")
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	//defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return database
}
