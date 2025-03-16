package rating_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingUsecaseI interface {
	RateProduct(r entities.Rating) error
	GetRates(productid primitive.ObjectID) []entities.RatingOut
	GetRating(productid, userid primitive.ObjectID) (float64, error)
}
