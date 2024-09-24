package models

import (
	"errors"

	// Import the models
	"datahandler_go/models/samples"
)

// Model is a placeholder for the type of your models. Replace with actual type.
type Model interface{}

// Define the row data models
var rowsDataModels = map[string]Model{
	"postgres_sample": &samples.Postgres_Sample{},
}

// Define the snapshot data models
var snapshotDataModels = map[string]Model{
	"mongo_sample": &samples.Mongo_Sample{},
}

// Merge row and snapshot data models
var models = mergeModels(rowsDataModels, snapshotDataModels)

// Function to merge maps
func mergeModels(maps ...map[string]Model) map[string]Model {
	merged := make(map[string]Model)
	for _, m := range maps {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

// Check if provided model name is a rows data model
func IsRowsDataModel(modelName string) bool {
	_, exists := rowsDataModels[modelName]
	return exists
}

// Check if provided model name is a snapshot data model
func IsSnapshotDataModel(modelName string) bool {
	_, exists := snapshotDataModels[modelName]
	return exists
}

// Retrieve the Model from the list of user-defined models
func GetModel(modelName string) (Model, error) {
	model, exists := models[modelName]
	if !exists {
		return nil, errors.New("model does not exist")
	}
	return model, nil
}
