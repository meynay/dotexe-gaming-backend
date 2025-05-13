package fave_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/category_rep"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/user_rep"
)

type FaveUsecase struct {
	userrep     *user_rep.UserRepository
	productrep  *product_rep.ProductRep
	categoryrep *category_rep.CategoryRep
}

func NewFaveUsecase(ur *user_rep.UserRepository, pr *product_rep.ProductRep, cr *category_rep.CategoryRep) *FaveUsecase {
	return &FaveUsecase{userrep: ur, productrep: pr, categoryrep: cr}
}

func (fu *FaveUsecase) FaveProduct(userid, productid uint) error {
	return fu.userrep.AddToFaves(productid, userid)
}

func (fu *FaveUsecase) UnfaveProduct(userid, productid uint) error {
	return fu.userrep.DeleteFromFaves(productid, userid)
}

func (fu *FaveUsecase) CheckFave(userid, productid uint) error {
	return fu.userrep.CheckFave(productid, userid)
}

func (fu *FaveUsecase) GetFaves(userid uint) []entities.ProductLess {
	pr := []entities.ProductLess{}
	ids := fu.userrep.GetFaves(userid)
	for _, id := range ids {
		product, _ := fu.productrep.GetProduct(id)
		category, _ := fu.categoryrep.GetCategory(product.CategoryID)
		pr = append(pr, entities.ProductLess{
			ID:          product.ID,
			Name:        product.Name,
			Image:       product.Image,
			Description: product.Description,
			Price:       product.Price,
			Category:    category.Name,
			Off:         product.Off,
			Rating:      product.Rating,
			RateCount:   product.RateCount,
		})
	}
	return pr
}
