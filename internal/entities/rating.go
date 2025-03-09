package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Rate      float64            `json:"rate" bson:"rate"`
	Review    string             `json:"review" bson:"review"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Likes     int                `json:"likes" bson:"likes"`
	Dislikes  int                `json:"dislikes" bson:"dislikes"`
}

type RatingOut struct {
	ID        primitive.ObjectID `json:"_id"`
	Username  string             `json:"user"`
	Rate      float64            `json:"rate"`
	Review    string             `json:"review"`
	CreatedAt time.Time          `json:"created_at"`
	Likes     int                `json:"likes"`
	Dislikes  int                `json:"dislikes"`
}
