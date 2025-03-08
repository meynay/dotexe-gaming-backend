package product_usecase

import "store/internal/entities"

type ProductUseCaseI interface {
	GetProduct(ID string) (entities.Product, error)
	GetProducts() ([]entities.ProductLess, error)
	FilterProducts(f entities.Filter) ([]entities.ProductLess, error)
	AddRatingToProduct(rate float64, ID string) error
	ChangeProductRating(oldRate, newRate float64, ID string) error
	AddPurchaseCount(count int, ID string) error
}
