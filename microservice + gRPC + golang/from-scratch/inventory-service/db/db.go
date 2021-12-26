package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDbConnection() *mongo.Client {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("INVENTORY_MYSQL_URL")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", port))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}

	return client
}
