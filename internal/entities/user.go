package entities

import "time"

type User struct {
	ID        string    `json:"id"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}
