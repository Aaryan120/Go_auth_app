package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var UsersCollection *mongo.Collection

func Connect() {
	// MongoDB connection string (connecting with database)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONNECTION_STRING"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Testing the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Setting the global variables
	Client = client
	UsersCollection = client.Database(os.Getenv("DB_NAME")).Collection("users")

	log.Println("Connected to MongoDB")
}
