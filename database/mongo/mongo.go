package mongo

import (
	"context"
	"datahandler_go/helpers"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectDb() {
	if IsDbConnected() {
		return
	}

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err error
	// Use mongo-db instead of localhost because this application is running within Docker
	uri := fmt.Sprintf("mongodb://%s:%s@mongo-db:%s/%s?authSource=admin",
		helpers.EnvVariable("MONGO_DB_USER"),
		helpers.EnvVariable("MONGO_DB_PASSWORD"),
		helpers.EnvVariable("MONGO_PORT"),
		helpers.EnvVariable("MONGO_DB_NAME"))
	client, err = mongo.Connect(ctx, options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(1).
		SetSocketTimeout(60*time.Second))

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
}

func DisconnectDb() {
	if !IsDbConnected() {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Disconnect(ctx)
	if err != nil {
		log.Println("Error closing database connection:", err)
	} else {
		log.Println("Database connection closed successfully.")
	}
}

func IsDbConnected() bool {
	if client == nil {
		return false
	}

	// Check MongoDB connection by pinging the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := client.Ping(ctx, readpref.Primary())
	return err == nil
}
