package entities

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	UserID    uint    `gorm:"index;not null" json:"-"`
	User      User    `gorm:"foreignKey:UserID" json:"-"`
	ProductID uint    `json:"product_id" bson:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Count     int     `json:"count" bson:"count"`
	Price     int     `json:"price" bson:"price"`
	Off       float64 `json:"off" bson:"off"`
}

type CartItem struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Image       string  `json:"image"`
	Count       int     `json:"count"`
	Price       int     `json:"price"`
	Off         float64 `json:"off"`
}

type Invoice struct {
	gorm.Model
	UserID        uint    `gorm:"index;not null" json:"user_id"`
	User          User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
	OrderStatus   int     `gorm:"type:smallint;default:0;not null" json:"order_status"`
	Items         []Item  `gorm:"type:jsonb;not null" json:"items"`
	DeliveryPrice int     `gorm:"type:integer;not null;default:0" json:"delivery_price"`
	CouponApplied bool    `gorm:"default:false;not null" json:"coupon_applied"`
	CouponOff     float64 `gorm:"type:numeric(5,2);default:0.0" json:"coupon_off"`
	CouponPrice   int     `gorm:"type:integer;default:0" json:"coupon_price"`
	TotalPrice    int     `gorm:"type:integer;not null" json:"total"`
}

type InvoiceFilter struct {
	Status      int
	From        time.Time
	To          time.Time
	CountToShow int
	Page        int
}

const TimeLayout = "2006-01-02"

const (
	Processing = iota
	GettingReady
	ReadyToPost
	Posted
	All
)
