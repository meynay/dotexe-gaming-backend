package product_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductRepI interface {
	AddProduct(p entities.Product) error
	GetProduct(ID primitive.ObjectID) (entities.Product, error)
	GetProducts() ([]entities.Product, error)
	EditProduct(p entities.Product) error
	DeleteProduct(ID primitive.ObjectID) error
	AddViewToProduct(ID primitive.ObjectID) error
	AddRatingToProduct(rate float64, ID primitive.ObjectID) error
	ChangeProductRating(oldRate, newRate float64, ID primitive.ObjectID) error
	AddPurchaseCount(count int, ID primitive.ObjectID) error
}
