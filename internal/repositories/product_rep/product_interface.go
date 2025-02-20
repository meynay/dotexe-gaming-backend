package product_rep

import "store/internal/entities"

type ProductRepI interface {
	AddProduct(p entities.Product) error
	GetProduct(ID string) (entities.Product, error)
	GetProducts() ([]entities.Product, error)
	EditProduct(p entities.Product) error
	DeleteProduct(ID string) error
	AddViewToProduct(ID string) error
	AddRatingToProduct(rate float64, ID string) error
	ChangeProductRating(oldRate, newRate float64, ID string) error
	AddPurchaseCount(count int, ID string) error
	AddCategory(c entities.Category) error
	EditCategory(c entities.Category) error
	DeleteCategory(ID string) error
	GetCategoryName(ID string) (string, error)
	GetParents(ID string) []string
}
