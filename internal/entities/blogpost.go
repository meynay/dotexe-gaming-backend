package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Content     string             `json:"content" bson:"content"`
	Category_id primitive.ObjectID `json:"category_id" bson:"categpry_id"`
	Image       string             `json:"image" bson:"image"`
	Author_id   primitive.ObjectID `json:"author_id" bson:"author_id"`
	Tags        []string           `json:"tags" bson:"tags"`
	Likes       int                `json:"likes" bson:"likes"`
	Dislikes    int                `json:"dislikes" bson:"dislikes"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type BPComment struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	BlogPostID primitive.ObjectID `json:"blogpost_id" bson:"blogpost_id"`
	Replier    string             `json:"replier" bson:"replier"`
	IsAdmin    bool               `json:"is_admin" bson:"is_admin"`
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
	Category    string             `json:"category"`
	Tags        []string           `json:"tags"`
	Likes       int                `json:"likes"`
	Dislikes    int                `json:"dislikes"`
	UpdatedAt   time.Time          `json:"updated_at"`
	CreatedAt   time.Time          `json:"created_at"`
}

type MiniBP struct {
	ID        primitive.ObjectID `json:"_id"`
	Title     string             `json:"title"`
	Category  string             `json:"category"`
	Image     string             `json:"image"`
	Author    string             `json:"author"`
	Likes     int                `json:"likes"`
	Dislikes  int                `json:"dislikes"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type BPFilter struct {
	Page       int                  `json:"page"`
	Count      int                  `json:"count"`
	Categories []primitive.ObjectID `json:"categories"`
	Query      string               `json:"query"`
}
