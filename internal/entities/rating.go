package entities

import "time"

type Rating struct {
	ID        string    `json:"_id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Rate      float64   `json:"rate"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

type RatingOut struct {
	ID        string    `json:"_id"`
	Username  string    `json:"user"`
	Rate      float64   `json:"rate"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}
