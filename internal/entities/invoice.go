package entities

import "time"

type Item struct {
	ProductID string  `json:"product_id"`
	Count     int     `json:"count"`
	Price     int     `json:"price"`
	Off       float64 `json:"off"`
}

type Invoice struct {
	ID            string    `json:"_id"`
	UserID        string    `json:"user_id"`
	InvoiceDate   time.Time `json:"date"`
	OrderStatus   int       `json:"order_status"`
	Items         []Item    `json:"items"`
	DeliveryPrice int       `json:"delivery_price"`
	CouponApplied bool      `json:"coupon_applied"`
	CouponOff     float64   `json:"coupon_off"`
	CouponPrice   int       `json:"coupon_price"`
	TotalPrice    int       `json:"total"`
}

type InvoiceFilter struct {
	Status      int
	From        time.Time
	To          time.Time
	CountToShow int
}

const (
	Processing = iota
	GettingReady
	ReadyToPost
	Posted
)
