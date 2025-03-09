package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Image     string             `json:"image" bson:"image"`
	Author_id primitive.ObjectID `json:"author_id" bson:"author_id"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type BPComment struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	BlogPostID primitive.ObjectID `json:"blogpost_id" bson:"blogpost_id"`
	Replier    string             `json:"replier" bson:"replier"`
	Comment    string             `json:"comment" bson:"comment"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
}

type BlogPostR struct {
	ID          primitive.ObjectID `json:"_id"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	Image       string             `json:"image"`
	Author      string             `json:"author"`
	AuthorImage string             `json:"author_image"`
	CreatedAt   time.Time          `json:"created_at"`
}
