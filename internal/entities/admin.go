package entities

import "time"

type Admin struct {
	ID        string `json:"_id"`
	Phone     string `json:"phone_number"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Image     string `json:"image"`
	Bio       string `json:"bio"`
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
