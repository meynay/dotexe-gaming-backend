package entities

import (
	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Type      int    `gorm:"type:smallint;not null;index" json:"activity_type"`
	Payload   string `gorm:"type:jsonb;not null" json:"payload"`
	IPAddress string `gorm:"type:varchar(45)" json:"ip_address"`
	UserAgent string `gorm:"type:text" json:"user_agent"`
	UserID    uint   `gorm:"index" json:"user_id"`
}

const (
	AddProductActivity = iota
	EditProductActivity
	DeleteProductActivity
	UserSignupActivity
	AddOrderActivity
	ChangeOrderStatusActivity
)
