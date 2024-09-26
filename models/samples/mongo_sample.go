package samples

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Define the model struct
type Mongo_Sample struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
	Label     string             `bson:"label" json:"label"`
	Value     float64            `bson:"value" json:"value"`
}
