package product_usecase

import "store/internal/entities"

type ProductUseCaseI interface {
	AddProduct(p entities.Product) error
	GetProduct(ID string) (entities.Product, error)
	GetProducts() ([]entities.ProductLess, error)
	EditProduct(p entities.Product) error
	FilterProducts(f entities.Filter) ([]entities.ProductLess, error)
	DeleteProduct(ID string) error
	AddRatingToProduct(rate float64, ID string) error
	ChangeProductRating(oldRate, newRate float64, ID string) error
	AddPurchaseCount(count int, ID string) error
	AddCategory(c entities.Category) error
}
