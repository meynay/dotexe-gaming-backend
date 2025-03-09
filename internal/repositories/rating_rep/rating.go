package rating_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RatingRep struct {
	rep *mongo.Collection
}

func NewRatingRep(rrep *mongo.Collection) *RatingRep {
	return &RatingRep{rep: rrep}
}
func (rr *RatingRep) AddRating(rating entities.Rating) error {
	_, err := rr.rep.InsertOne(context.TODO(), rating)
	if err != nil {
		return fmt.Errorf("couldn't insert rate")
	}
	return nil
}

func (rr *RatingRep) ChangeRating(rating entities.Rating) error {
	_, err := rr.rep.DeleteOne(context.TODO(), bson.M{"user_id": rating.UserID, "product_id": rating.ProductID})
	if err != nil {
		return fmt.Errorf("can't delete existing rate")
	}
	_, err = rr.rep.InsertOne(context.TODO(), rating)
	if err != nil {
		return fmt.Errorf("couldn't insert rate")
	}
	return nil
}

func (rr *RatingRep) GetRating(userid, productid primitive.ObjectID) (entities.Rating, error) {
	res := rr.rep.FindOne(context.TODO(), bson.M{"user_id": userid, "product_id": productid})
	if res.Err() != nil {
		return entities.Rating{}, res.Err()
	}
	var r entities.Rating
	res.Decode(&r)
	return r, nil
}

func (rr *RatingRep) GetRatings(productid primitive.ObjectID) []entities.Rating {
	var rates []entities.Rating
	res, err := rr.rep.Find(context.TODO(), bson.M{"product_id": productid})
	if err != nil {
		return rates
	}
	res.All(context.TODO(), &rates)
	return rates
}
