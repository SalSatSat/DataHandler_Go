package routes

import (
	"context"
	"datahandler_go/database/mongo"
	"datahandler_go/helpers"
	"datahandler_go/models"
	"encoding/csv"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Function to send records as CSV
func sendCSV(c *fiber.Ctx, records []bson.M) error {
	c.Set("Content-Type", "text/csv")
	writer := csv.NewWriter(c.Response().BodyWriter())

	if len(records) == 0 {
		return nil // No records to write
	}

	// Write header dynamically from the first record
	var header []string
	for key := range records[0] {
		header = append(header, key)
	}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, record := range records {
		var row []string
		for _, key := range header { // Use the dynamic header for writing rows
			row = append(row, fmt.Sprintf("%v", record[key]))
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func SetupRoutes(app *fiber.App) {
	app.Get("/*", func(c *fiber.Ctx) error {
		modelToQuery := strings.TrimPrefix(c.Path(), "/")

		// Reject file requests, e.g. favicon.ico
		re := regexp.MustCompile(`\..*`)
		if re.MatchString(modelToQuery) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":       "Invalid request",
				"status_code": fiber.StatusBadRequest,
			})
		}

		sort := c.Query("sort", "")
		skip := c.QueryInt("skip", 0)
		limit := c.QueryInt("limit", 100)

		// Connect to MongoDB
		mongo.ConnectDb()
		defer mongo.DisconnectDb()

		if mongo.IsDbConnected() {
			fmt.Printf("Mongo successfully connected\n")

			collection := mongo.Client.Database(helpers.EnvVariable("MONGO_DB_NAME")).Collection(modelToQuery)

			// Apply query based on request parameters
			var records []bson.M
			var err error

			if models.IsRowsDataModel(modelToQuery) {
				findOptions := options.Find().SetSkip(int64(skip)).SetLimit(int64(limit))

				// Validate sort parameter
				if sort != "" {
					findOptions.SetSort(bson.D{{Key: sort, Value: 1}}) // Only set sort if it's not empty
				} else {
					// Optionally log or handle the case where sort is empty
					fmt.Println("Warning: sort parameter is empty; default sorting will be applied.")
				}

				cursor, err := collection.Find(context.Background(), bson.M{}, findOptions)
				if err != nil {
					return err
				}
				if err := cursor.All(context.Background(), &records); err != nil {
					return err
				}

				// Convert records to CSV
				if err := sendCSV(c, records); err != nil {
					return err
				}
			} else {
				// Handle the case for a single document
				var record bson.M
				err = collection.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.D{{Key: "timestamp", Value: -1}})).Decode(&record)
				if err != nil {
					if err == mongoDriver.ErrNoDocuments {
						return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
							"error":       "No document found",
							"status_code": fiber.StatusNotFound,
						})
					}
					return err
				}
				// Send the single record as JSON response
				return c.JSON(record)
			}

			return nil
		} else {
			log.Println("Mongo failed to connect")
		}
		return nil
	})
}
