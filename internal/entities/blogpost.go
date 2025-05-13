package entities

import (
	"time"

	"gorm.io/gorm"
)

type BlogPost struct {
	gorm.Model

	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	Image   string `gorm:"type:text" json:"image"`

	CategoryID uint     `gorm:"index" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"`

	AuthorID uint  `gorm:"index" json:"author_id"`
	Author   Admin `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"author"`

	Tags []string `gorm:"type:jsonb" json:"tags"`

	Likes    int `gorm:"default:0" json:"likes"`
	Dislikes int `gorm:"default:0" json:"dislikes"`
}

type BPComment struct {
	gorm.Model

	BlogPostID uint     `gorm:"index" json:"blogpost_id"`
	BlogPost   BlogPost `gorm:"foreignKey:BlogPostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"blog_post"`

	ReplierID uint   `gorm:"index" json:"replier_id"`
	IsAdmin   bool   `gorm:"default:false" json:"is_admin"`
	Comment   string `gorm:"type:text;not null" json:"comment"`
}

type BlogPostR struct {
	ID          uint      `json:"id"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	Author      string    `json:"author"`
	AuthorImage string    `json:"author_image"`
	Category    string    `json:"category"`
	Tags        []string  `json:"tags"`
	Likes       int       `json:"likes"`
	Dislikes    int       `json:"dislikes"`
}

type MiniBP struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Category  string    `json:"category"`
	Image     string    `json:"image"`
	Author    string    `json:"author"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

type BPFilter struct {
	Page       int    `json:"page"`
	Count      int    `json:"count"`
	Categories []uint `json:"categories"`
	Query      string `json:"query"`
}
