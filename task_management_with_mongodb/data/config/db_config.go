package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const connectionStrings = "mongodb+srv://mente:123mongodb1234@cluster0.evfwybz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
// const dbName = "test"
// const collectionName = "tasks"

var Collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables.")
	}
	connectionStrings := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	if connectionStrings == "" || dbName == "" || collectionName == "" {
		log.Fatal("Environment variables not set")
	}
	
	clientOptions := options.Client().ApplyURI(connectionStrings)

	client,err:=mongo.Connect(context.TODO(),clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Connection to mongodb is ready")
}
