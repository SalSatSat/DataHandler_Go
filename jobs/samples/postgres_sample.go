package samples

import (
	"context"
	"datahandler_go/database/mongo"
	"datahandler_go/database/postgres"
	"datahandler_go/helpers"
	"datahandler_go/models/samples"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var records []samples.Postgres_Sample
var updates = make(map[int]bool)

func tryEndJob(startTime time.Time) {
	allUpdated := true
	for _, updated := range updates {
		if !updated {
			allUpdated = false
			break
		}
	}

	if allUpdated {
		jobTime := time.Since(startTime).Seconds()
		log.Printf("postgres_sample job completed in %.2f seconds.\n", jobTime)
	}
}

func Postgres_Sample_Job() {
	startTime := time.Now()
	log.Println("Starting postgres_sample job...")

	// Connect to PostgreSQL and fetch data
	postgres.ConnectDb()
	defer postgres.DisconnectDb()

	postgres.DB.Db.AutoMigrate(&samples.Postgres_Sample{})

	if postgres.IsDbConnected() {
		log.Println("Postgres successfully connected")

		// Fetch data from PostgreSQL table into a slice
		records = []samples.Postgres_Sample{}
		if err := postgres.DB.Db.Find(&records).Error; err != nil {
			log.Fatalf("Error fetching data from Postgres: %v\n", err)
		}

		log.Printf("Fetched %d records from Postgres\n", len(records))
	} else {
		log.Println("Postgres failed to connect")
		return
	}

	// Connect to MongoDB
	mongo.ConnectDb()
	defer mongo.DisconnectDb()

	if mongo.IsDbConnected() {
		log.Println("Mongo successfully connected")

		// Define MongoDB collection
		collection := mongo.Client.Database(helpers.EnvVariable("MONGO_DB_NAME")).Collection("postgres_sample")

		// Process each row
		for _, record := range records {
			filter := bson.M{"label": record.Label} // Use a unique field for the filter
			update := bson.M{
				"$set": bson.M{
					"UpdatedAt": record.UpdatedAt,
					"label":     record.Label,
					"value":     record.Value,
				},
			}

			updateOptions := options.FindOneAndUpdate().SetUpsert(true) // Upsert option

			var updatedSample samples.Mongo_Sample
			err := collection.FindOneAndUpdate(context.TODO(), filter, update, updateOptions).Decode(&updatedSample)

			if err != nil {
				if err == mongoDriver.ErrNoDocuments {
					// Document was not found, but we used upsert, so no error
					// You may log that the document was created if needed
					log.Printf("Inserted new document: %+v\n", record)
				} else {
					log.Fatal(err)
				}
			} else {
				updates[int(record.ID)] = true
				log.Printf("Updated document: %+v\n", updatedSample)
			}
		}

		tryEndJob(startTime)
	} else {
		log.Println("Mongo failed to connect")
	}
}
