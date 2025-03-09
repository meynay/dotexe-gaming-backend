package rating_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingRepI interface {
	AddRating(rating entities.Rating) error
	ChangeRating(rating entities.Rating) error
	GetRating(userid, productid primitive.ObjectID) (entities.Rating, error)
	GetRatings(productid primitive.ObjectID) []entities.Rating
}
