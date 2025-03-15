package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Parent    primitive.ObjectID `json:"parent" bson:"parent"`
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	IsAdmin   bool               `json:"is_admin" bson:"is_admin"`
	Comment   string             `json:"comment" bson:"comment"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type CommentOut struct {
	ID        primitive.ObjectID `json:"_id"`
	Parent    primitive.ObjectID `json:"parent"`
	User      string             `json:"user"`
	Comment   string             `json:"comment"`
	CreatedAt time.Time          `json:"created_at"`
}
