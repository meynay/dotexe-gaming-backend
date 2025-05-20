package rating_rep

import (
	"fmt"
	"store/internal/entities"

	"gorm.io/gorm"
)

type RatingRep struct {
	rep *gorm.DB
}

func NewRatingRep(rrep *gorm.DB) *RatingRep {
	return &RatingRep{rep: rrep}
}
func (rr *RatingRep) AddRating(rating entities.Rating) error {
	tx := rr.rep.Create(&rating)
	if tx.Error != nil {
		return fmt.Errorf("couldn't insert rate")
	}
	return nil
}

func (rr *RatingRep) ChangeRating(rating entities.Rating) error {
	tx := rr.rep.Delete(&entities.Rating{}, entities.Rating{UserID: rating.UserID, ProductID: rating.ProductID})
	if tx.Error != nil {
		return fmt.Errorf("can't delete existing rate")
	}
	tx = rr.rep.Create(&rating)
	if tx.Error != nil {
		return fmt.Errorf("couldn't insert rate")
	}
	return nil
}

func (rr *RatingRep) GetRating(userid, productid uint) (entities.Rating, error) {
	var r entities.Rating
	res := rr.rep.Where("user_id = ? AND product_id = ?", userid, productid).First(&r)
	if res.Error != nil {
		return entities.Rating{}, res.Error
	}
	return r, nil
}

func (rr *RatingRep) GetRatings(productid uint) []entities.Rating {
	var rates []entities.Rating
	tx := rr.rep.Where("product_id = ?", productid).Find(&rates)
	if tx.Error != nil {
		return []entities.Rating{}
	}
	return rates
}
