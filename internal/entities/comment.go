package entities

import "time"

type Comment struct {
	ID        string    `json:"_id"`
	Parent    string    `json:"parent"`
	ProductID string    `json:"product_id"`
	UserID    string    `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentOut struct {
	ID        string    `json:"_id"`
	Parent    string    `json:"parent"`
	User      string    `json:"user"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
