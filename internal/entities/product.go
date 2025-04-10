package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Image    string             `json:"image" bson:"image"`
	ParentID primitive.ObjectID `json:"parent_id" bson:"parent_id"`
}
type Product struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name          string             `json:"name" bson:"name"`
	Image         string             `json:"image" bson:"image"`
	Images        []string           `json:"images" bson:"images"`
	Description   string             `json:"description" bson:"description"`
	Price         int                `json:"price" bson:"price"`
	Stock         int                `json:"stock" bson:"stock"`
	Off           float64            `json:"off" bson:"off"`
	Info          map[string]string  `json:"info" bson:"info"`
	CategoryID    primitive.ObjectID `json:"category_id" bson:"category_id"`
	AddedAt       time.Time          `json:"time_added" bson:"time_added"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
	Rating        float64            `json:"rating" bson:"rating"`
	RateCount     int                `json:"rate_count" bson:"rate_count"`
	Views         int                `json:"views" bson:"views"`
	PurchaseCount int                `json:"purchase_count" bson:"purchase_count"`
	Tags          []string           `json:"tags" bson:"tags"`
}

type ProductLess struct {
	ID          primitive.ObjectID `json:"_id"`
	Name        string             `json:"name"`
	Image       string             `json:"image"`
	Description string             `json:"description"`
	Price       int                `json:"price"`
	Category    string             `json:"category"`
	Off         float64            `json:"off"`
	Rating      float64            `json:"rating"`
	RateCount   int                `json:"rate_count"`
}

type Filter struct {
	Query         string             `json:"query"`
	CategoryID    primitive.ObjectID `json:"category_id"`
	Page          int                `json:"page"`
	NumberOfItems int                `json:"number_of_items"`
	Order         int                `json:"order"`
}

type PScore struct {
	Pr    Product
	Score float64
}

const (
	CheapToExpensive = iota
	ExpensiveToCheap
	MostOffToLeast
	Newest
	MostViewed
	MostPurchased
	MostRelevant
	MostRate
)
