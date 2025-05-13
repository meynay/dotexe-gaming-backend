package entities

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model

	Username  string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Phone     string `gorm:"type:varchar(20);unique;not null" json:"phone_number"`
	Password  string `gorm:"type:varchar(100);not null" json:"-"`
	FirstName string `gorm:"type:varchar(50);not null" json:"firstname"`
	LastName  string `gorm:"type:varchar(50);not null" json:"lastname"`
	Image     string `gorm:"type:text" json:"image"`
	Bio       string `gorm:"type:text" json:"bio"`
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
