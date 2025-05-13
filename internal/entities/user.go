package entities

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID      uint   `gorm:"index;not null" json:"-"`
	Province    string `json:"province" gorm:"type:varchar(100);not null"`
	City        string `json:"city" gorm:"type:varchar(100);not null"`
	Addr        string `json:"address" gorm:"type:text;not null"`
	PhoneNumber string `json:"phone_number" gorm:"type:varchar(20);not null"`
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	User        User   `gorm:"foreignKey:UserID" json:"-"`
}

type User struct {
	gorm.Model
	Phone        string    `gorm:"type:varchar(20);uniqueIndex;not null" json:"phone_number"`
	Password     string    `gorm:"type:varchar(100);not null" json:"-"`
	Email        string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	RefreshToken string    `gorm:"type:text;index" json:"-"`
	FirstName    string    `gorm:"type:varchar(50);not null" json:"firstname"`
	LastName     string    `gorm:"type:varchar(50);not null" json:"lastname"`
	Addresses    []Address `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"addresses,omitempty"`
	Faves        []uint    `gorm:"type:jsonb" json:"faves"`
	Cart         []Item    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"cart,omitempty"`
	LastLoginAt  time.Time `gorm:"index" json:"-"`
	IsVerified   bool      `gorm:"default:false" json:"is_verified"`
}
