package entities

import "time"

type Category struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Image    string `json:"image"`
	ParentID string `json:"parent_id"`
}

type Product struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Image       string            `json:"image"`
	Images      []string          `json:"images"`
	Description string            `json:"description"`
	Price       int               `json:"price"`
	Stock       int               `json:"stock"`
	Off         float64           `json:"off"`
	Info        map[string]string `json:"info"`
	CategoryID  string            `json:"category_id"`
	AddedAt     time.Time         `json:"time_added"`
}

type ProductLess struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Image    string  `json:"image"`
	Price    int     `json:"price"`
	Category string  `json:"category"`
	Off      float64 `json:"off"`
}

type Filter struct {
	Query         string `json:"query"`
	CategoryID    string `json:"category_id"`
	Page          int    `json:"page"`
	NumberOfItems int    `json:"number_of_items"`
	Order         int    `json:"order"`
}

type PScore struct {
	Pr    Product
	Score float64
}

const (
	CheapToExpensive = iota + 100
	ExpensiveToCheap
	MostOffToLeast
	Newest
	MostViewed
	MostPurchased
	MostRelevant
)
