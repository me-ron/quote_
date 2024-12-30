package config

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() error {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables.")
	}

	uri := os.Getenv("DB_URL")
	if uri == "" {
		return errors.New("missing DB_URL environment variable, please set it")
	}

	// Create MongoDB client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Ping to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	DB = client.Database("quoteDB")
	log.Println("Successfully connected to MongoDB and initialized the database.")
	return nil
}
