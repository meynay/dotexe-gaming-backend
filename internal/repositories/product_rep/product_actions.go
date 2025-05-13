package product_rep

import (
	"fmt"
	"store/internal/entities"
)

func (pr *ProductRep) AddRatingToProduct(rate float64, ID uint) error {
	product := entities.Product{}
	res := pr.prdb.First(&product, ID)
	if res.Error != nil {
		return fmt.Errorf("error getting product, %v", res.Error)
	}
	product.Rating *= float64(product.RateCount)
	product.RateCount++
	product.Rating += rate
	product.Rating = product.Rating / float64(product.RateCount)
	tx := pr.prdb.Save(product)
	if tx.Error != nil {
		return fmt.Errorf("error occured saving to db, %v", tx.Error)
	}
	return nil
}

func (pr *ProductRep) ChangeProductRating(oldRate, newRate float64, ID uint) error {
	product := entities.Product{}
	res := pr.prdb.First(&product, ID)
	if res.Error != nil {
		return fmt.Errorf("error getting product, %v", res.Error)
	}
	product.Rating *= float64(product.RateCount)
	product.Rating += newRate - oldRate
	product.Rating /= float64(product.RateCount)
	tx := pr.prdb.Save(product)
	if tx.Error != nil {
		return fmt.Errorf("error occured saving to db, %v", tx.Error)
	}
	return nil
}

func (pr *ProductRep) AddPurchaseCount(count int, ID uint) error {
	product := entities.Product{}
	res := pr.prdb.First(&product, ID)
	if res.Error != nil {
		return fmt.Errorf("error getting product, %v", res.Error)
	}
	product.PurchaseCount += count
	tx := pr.prdb.Save(product)
	if tx.Error != nil {
		return fmt.Errorf("error occured saving to db, %v", tx.Error)
	}
	return nil
}

func (pr *ProductRep) AddViewToProduct(ID uint) error {
	product := entities.Product{}
	res := pr.prdb.First(&product, ID)
	if res.Error != nil {
		return fmt.Errorf("error getting product, %v", res.Error)
	}
	product.Views++
	tx := pr.prdb.Save(product)
	if tx.Error != nil {
		return fmt.Errorf("error occured saving to db, %v", tx.Error)
	}
	return nil
}
