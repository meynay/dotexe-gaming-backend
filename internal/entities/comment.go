package entities

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model

	ParentID  *uint    `gorm:"default:null;index" json:"parent_id"`
	Parent    *Comment `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"parent,omitempty"`
	ProductID uint     `gorm:"index;not null" json:"product_id"`
	Product   Product  `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	UserID    uint     `gorm:"index;not null" json:"user_id"`
	User      User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user,omitempty"`
	IsAdmin   bool     `gorm:"default:false;not null" json:"is_admin"`
	Comment   string   `gorm:"type:text;not null" json:"comment"`
}

type CommentOut struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Parent    uint      `json:"parent"`
	User      string    `json:"user"`
	Comment   string    `json:"comment"`
}
