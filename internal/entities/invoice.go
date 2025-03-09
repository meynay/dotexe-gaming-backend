package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ProductID primitive.ObjectID `json:"product_id" bson:"product_id"`
	Count     int                `json:"count" bson:"count"`
	Price     int                `json:"price" bson:"price"`
	Off       float64            `json:"off" bson:"off"`
}

type Invoice struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	UserID        primitive.ObjectID `json:"user_id" bson:"user_id"`
	InvoiceDate   time.Time          `json:"date" bson:"date"`
	OrderStatus   int                `json:"order_status" bson:"order_status"`
	Items         []Item             `json:"items" bson:"items"`
	DeliveryPrice int                `json:"delivery_price" bson:"delivery_price"`
	CouponApplied bool               `json:"coupon_applied" bson:"coupon_applied"`
	CouponOff     float64            `json:"coupon_off" bson:"coupon_off"`
	CouponPrice   int                `json:"coupon_price" bson:"coupon_price"`
	TotalPrice    int                `json:"total" bson:"total"`
}

type InvoiceFilter struct {
	Status      int
	From        time.Time
	To          time.Time
	CountToShow int
	Page        int
}

const (
	Processing = iota
	GettingReady
	ReadyToPost
	Posted
	All
)
