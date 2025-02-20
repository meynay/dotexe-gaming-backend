package entities

import "time"

type Rating struct {
	ID        string    `json:"_id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Rate      float64   `json:"rate"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

type RatingOut struct {
	ID        string    `json:"_id"`
	Username  string    `json:"user"`
	Rate      float64   `json:"rate"`
	Review    string    `json:"review"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
}

type Comment struct {
	ID        string    `json:"_id"`
	Parent    string    `json:"parent"`
	ProductID string    `json:"product_id"`
	UserID    string    `json:"user_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentOut struct {
	ID        string    `json:"_id"`
	Parent    string    `json:"parent"`
	User      string    `json:"user"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

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
	Items         []Item    `json:"items"`
	DeliveryPrice int       `json:"delivery_price"`
	CouponApplied bool      `json:"coupon_applied"`
	CouponOff     float64   `json:"coupon_off"`
	CouponPrice   int       `json:"coupon_price"`
	TotalPrice    int       `json:"total"`
}
