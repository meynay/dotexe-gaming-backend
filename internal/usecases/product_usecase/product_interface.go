package product_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCaseI interface {
	GetProduct(ID primitive.ObjectID) (entities.Product, error)
	GetProducts() ([]entities.ProductLess, error)
	FilterProducts(f entities.Filter) ([]entities.ProductLess, error)
	AddRatingToProduct(rate float64, ID primitive.ObjectID) error
	ChangeProductRating(oldRate, newRate float64, ID primitive.ObjectID) error
	AddPurchaseCount(count int, ID primitive.ObjectID) error
	GetCategories() []entities.Category
}
