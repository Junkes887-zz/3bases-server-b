package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Context struct {
	CTX context.Context
}

func (context Context) CreateConnectionMongo() *mongo.Client {
	MONGO_URL := os.Getenv("MONGO_URL")

	clientOptions := options.Client().ApplyURI(MONGO_URL)
	client, err := mongo.Connect(context.CTX, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected mongo...")
	return client
}
