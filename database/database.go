package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	connectionString = "enter_your_string"
	dbName           = "mydatabase"
)

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal("Error creating MongoDB client:", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	fmt.Println("MongoDB connection success")

	Client = client
	Collection = client.Database(dbName).Collection("users")
	fmt.Println("Collection instance is ready")
}

func DisconnectDB() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error disconnecting from MongoDB:", err)
	}
	fmt.Println("MongoDB connection closed")
}
