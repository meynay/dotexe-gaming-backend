package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Activities struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Type      string             `json:"activity_type" bson:"activity_type"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Payload   string             `json:"payload" bson:"payload"`
}
