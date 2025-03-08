package fave_usecase

import "store/internal/entities"

type FaveUsecaseI interface {
	FaveProduct(userid, productid string) error
	UnfaveProduct(userid, productid string) error
	CheckFave(userid, productid string) error
	GetFaves(userid string) []entities.ProductLess
}
