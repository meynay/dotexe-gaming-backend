package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username"`
	Phone     string             `json:"phone_number" bson:"phone_number"`
	Password  string             `json:"password" bson:"password"`
	FirstName string             `json:"firstname" bson:"firstname"`
	LastName  string             `json:"lastname" bson:"lastname"`
	Image     string             `json:"image" bson:"image"`
	Bio       string             `json:"bio" bson:"bio"`
}

type ChartFilter struct {
	From     time.Time
	To       time.Time
	ShowType int
}

const (
	OrdersCount = iota
	ItemsCount
	TotalPrice
)
