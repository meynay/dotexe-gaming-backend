package entities

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	UserID    uint     `gorm:"index;not null" json:"user_id"`
	User      User     `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	ProductID uint     `gorm:"index;not null" json:"product_id"`
	Product   Product  `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product,omitempty"`
	Rate      float64  `gorm:"type:numeric(2,1);not null;check:rate >= 1 AND rate <= 5" json:"rate"`
	Review    string   `gorm:"type:text" json:"review"`
	Likes     int      `gorm:"default:0" json:"likes"`
	Dislikes  int      `gorm:"default:0" json:"dislikes"`
	_         struct{} `gorm:"uniqueIndex:idx_user_product"`
}

type RatingOut struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"user"`
	Rate      float64   `json:"rate"`
	Review    string    `json:"review"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}
