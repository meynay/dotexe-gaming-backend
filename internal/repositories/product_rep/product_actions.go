package product_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
)

func (pr *ProductRep) AddRatingToProduct(rate float64, ID string) error {
	product := entities.Product{}
	res := pr.prdb.FindOne(context.TODO(), bson.M{"_id": ID})
	if res.Err() != nil {
		return fmt.Errorf("couldn't find product")
	}
	res.Decode(&product)
	product.Rating *= float64(product.RateCount)
	product.RateCount++
	product.Rating += rate
	product.Rating = product.Rating / float64(product.RateCount)
	pr.prdb.UpdateOne(context.TODO(), bson.M{"id": product.ID}, product)
	return nil
}

func (pr *ProductRep) ChangeProductRating(oldRate, newRate float64, ID string) error {
	product := entities.Product{}
	res := pr.prdb.FindOne(context.TODO(), bson.M{"_id": ID})
	if res.Err() != nil {
		return fmt.Errorf("couldn't find product")
	}
	res.Decode(&product)
	product.Rating *= float64(product.RateCount)
	product.Rating += newRate - oldRate
	product.Rating /= float64(product.RateCount)
	pr.prdb.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, product)
	return nil
}

func (pr *ProductRep) AddPurchaseCount(count int, ID string) error {
	product := entities.Product{}
	res := pr.prdb.FindOne(context.TODO(), bson.M{"_id": ID})
	if res.Err() != nil {
		return fmt.Errorf("couldn't find product")
	}
	res.Decode(&product)
	product.PurchaseCount += count
	pr.prdb.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, product)
	return nil
}

func (pr *ProductRep) AddViewToProduct(ID string) error {
	product := entities.Product{}
	res := pr.prdb.FindOne(context.TODO(), bson.M{"_id": ID})
	if res.Err() != nil {
		return fmt.Errorf("couldn't find product")
	}
	res.Decode(&product)
	product.Views++
	pr.prdb.UpdateOne(context.TODO(), bson.M{"_id": product.ID}, product)
	return nil
}
