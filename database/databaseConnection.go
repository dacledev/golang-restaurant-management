package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	MongoDB := "mongo://localhost:27017"
	fmt.Print(MongoDB)

	client, error := mongo.NewClient(options.Client().ApplyURI(MongoDB))

	if error != nil {
		log.Fatal(error)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	error = client.Connect(ctx)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}
