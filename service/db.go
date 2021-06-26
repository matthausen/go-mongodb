package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)

const (
	connectionString = ""
)

var collection *mongo.Collection

// Initialise the MongoDB client
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	database := os.Getenv("DATABASE")
	coll := os.Getenv("COLLECTION")

	clientOptions := options.Client().ApplyURI(connectionString)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	// Ping the DB
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Could not ping db: %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(database).Collection(coll)

	fmt.Println("Collection instance created!")
}


