package product_rep

import "store/internal/entities"

type ProductRepI interface {
	AddProduct(p entities.Product) error
	GetProduct(ID string) (entities.Product, error)
	GetProducts() ([]entities.Product, error)
	EditProduct(p entities.Product) error
	DeleteProduct(ID string) error
	AddCategory(c entities.Category) error
	GetCategoryName(ID string) (string, error)
	GetParents(ID string) []string
}
