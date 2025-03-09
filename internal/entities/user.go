package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	Phone     string               `json:"phone_number" bson:"phone_number"`
	Password  string               `json:"password" bson:"password"`
	Email     string               `json:"email" bson:"email"`
	FirstName string               `json:"firstname" bson:"firstname"`
	LastName  string               `json:"lastname" bson:"lastname"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	Faves     []primitive.ObjectID `json:"faves" bson:"faves"`
	Cart      []Item               `json:"cart" bson:"cart"`
}
