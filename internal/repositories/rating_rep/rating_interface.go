package rating_rep

import "store/internal/entities"

type RatingRepI interface {
	AddRating(rating entities.Rating) error
	ChangeRating(rating entities.Rating) error
	GetRating(userid, productid string) (entities.Rating, error)
	GetRatings(productid string) []entities.Rating
}
