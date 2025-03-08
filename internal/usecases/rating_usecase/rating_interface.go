package rating_usecase

import "store/internal/entities"

type RatingUsecaseI interface {
	RateProduct(r entities.Rating) error
	GetRates(productid string) []entities.RatingOut
}
