package entities

import "time"

type BlogPost struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     string    `json:"image"`
	Author_id string    `json:"authorid"`
	CreatedAt time.Time `json:"created_at"`
}

type BPComment struct {
	ID         string    `json:"id"`
	BlogPostID string    `json:"blogpost_id"`
	Replier    string    `json:"replier"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
}

type BlogPostR struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	Author      string    `json:"author"`
	AuthorImage string    `json:"author_image"`
	CreatedAt   time.Time `json:"created_at"`
}
