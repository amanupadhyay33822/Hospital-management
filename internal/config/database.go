package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDB() *mongo.Client {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/test"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")

	return Client
}

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection
	collection = client.Database("testdb").Collection(collectionName)
	return collection
}
