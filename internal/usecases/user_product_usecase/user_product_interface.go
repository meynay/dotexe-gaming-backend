package user_product_usecase

import "store/internal/entities"

type UserProductUsecaseI interface {
	FaveProduct(userid, productid string) error
	UnfaveProduct(userid, productid string) error
	CheckFave(userid, productid string) error
	GetFaves(userid string) []entities.ProductLess
	CommentOnProduct(c entities.Comment) error
	GetComments(productid string) []entities.CommentOut
	RateProduct(r entities.Rating) error
	GetRates(productid string) []entities.RatingOut
	AddToCart(productid, userid string) error
	IsInCart(productid, userid string) (int, error)
}
