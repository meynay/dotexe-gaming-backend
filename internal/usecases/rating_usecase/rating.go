package rating_usecase

import (
	"fmt"
	"store/internal/entities"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/rating_rep"
	"store/internal/repositories/user_rep"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingUsecase struct {
	ratingrep  *rating_rep.RatingRep
	productrep *product_rep.ProductRep
	userrep    *user_rep.UserRepository
}

func NewRatingUsecase(rr *rating_rep.RatingRep, pr *product_rep.ProductRep, ur *user_rep.UserRepository) *RatingUsecase {
	return &RatingUsecase{ratingrep: rr, productrep: pr, userrep: ur}
}

func (ru *RatingUsecase) RateProduct(r entities.Rating) error {
	or, err := ru.ratingrep.GetRating(r.UserID, r.ProductID)
	if err != nil {
		err = ru.ratingrep.AddRating(r)
		if err != nil {
			return fmt.Errorf("couldn't add rating")
		}
		return ru.productrep.AddRatingToProduct(r.Rate, r.ProductID)
	}
	err = ru.ratingrep.ChangeRating(r)
	if err != nil {
		return fmt.Errorf("couldn't change rating")
	}
	return ru.productrep.ChangeProductRating(or.Rate, r.Rate, r.ProductID)
}

func (ru *RatingUsecase) GetRates(productid primitive.ObjectID) []entities.RatingOut {
	rates := ru.ratingrep.GetRatings(productid)
	ratings := []entities.RatingOut{}
	for _, r := range rates {
		newrating := entities.RatingOut{
			ID:        r.ID,
			Rate:      r.Rate,
			Username:  ru.userrep.GetUsername(r.UserID),
			Review:    r.Review,
			CreatedAt: r.CreatedAt,
			Likes:     r.Likes,
			Dislikes:  r.Dislikes,
		}
		ratings = append(ratings, newrating)
	}
	return ratings
}
