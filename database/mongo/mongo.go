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

var Client *mongo.Client

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
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		helpers.EnvVariable("MONGO_DB_USER"),
		helpers.EnvVariable("MONGO_DB_PASSWORD"),
		helpers.EnvVariable("MONGO_HOST"),
		helpers.EnvVariable("MONGO_PORT"),
		helpers.EnvVariable("MONGO_DB_NAME"))
	Client, err = mongo.Connect(ctx, options.Client().
		ApplyURI(uri).
		SetMaxPoolSize(1).
		SetSocketTimeout(60*time.Second))

	if err != nil {
		log.Fatal("Failed to connect to Mongo database. \n", err)
	}
}

func DisconnectDb() {
	if !IsDbConnected() {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := Client.Disconnect(ctx)
	if err != nil {
		log.Println("Error closing Mongo database connection:", err)
	} else {
		log.Println("Mongo database connection closed successfully.")
	}
}

func IsDbConnected() bool {
	if Client == nil {
		return false
	}

	// Check MongoDB connection by pinging the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err := Client.Ping(ctx, readpref.Primary())
	return err == nil
}
