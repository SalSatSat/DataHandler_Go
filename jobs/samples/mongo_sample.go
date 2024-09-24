package samples

import (
	"context"
	"datahandler_go/database/mongo"
	"datahandler_go/helpers"
	"datahandler_go/models/samples"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createSampleData() []samples.Mongo_Sample {
	records := []samples.Mongo_Sample{
		{
			ID:        primitive.NewObjectID(),
			Timestamp: time.Now(),
			Label:     "Sample A",
			Value:     10.5,
		},
		{
			ID:        primitive.NewObjectID(),
			Timestamp: time.Now(),
			Label:     "Sample B",
			Value:     20.75,
		},
		{
			ID:        primitive.NewObjectID(),
			Timestamp: time.Now(),
			Label:     "Sample C",
			Value:     15.0,
		},
		{
			ID:        primitive.NewObjectID(),
			Timestamp: time.Now(),
			Label:     "Sample D",
			Value:     30.99,
		},
	}
	return records
}

func Mongo_Sample_Job() {
	startTime := time.Now()
	log.Println("Starting mongo_sample job...")

	// Connect to MongoDB
	mongo.ConnectDb()
	defer mongo.DisconnectDb()

	if mongo.IsDbConnected() {
		log.Println("Mongo successfully connected")

		// Define MongoDB collection
		collection := mongo.Client.Database(helpers.EnvVariable("MONGO_DB_NAME")).Collection("mongo_sample")

		// Create sample data
		records := createSampleData()

		// Upsert sample data into the collection
		for _, record := range records {
			filter := bson.M{"label": record.Label} // Unique filter based on Label
			update := bson.M{
				"$set": bson.M{
					"timestamp": record.Timestamp,
					"value":     record.Value,
				},
			}
			opts := options.FindOneAndUpdate().SetUpsert(true) // Upsert option

			var updatedSample samples.Mongo_Sample
			err := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedSample)

			if err != nil {
				if err == mongoDriver.ErrNoDocuments {
					// Document was not found, but we used upsert, so no error
					// You may log that the document was created if needed
					log.Printf("Inserted new document: %+v\n", record)
				} else {
					log.Fatal(err)
				}
			} else {
				log.Printf("Updated document: %+v\n", updatedSample)
			}
		}

		jobTime := time.Since(startTime).Seconds()
		log.Printf("mongo_sample job complete in %.2f seconds.\n", jobTime)
	} else {
		log.Println("Mongo failed to connect")
	}
}
