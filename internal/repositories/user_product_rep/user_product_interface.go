package user_product_rep

import (
	"store/internal/entities"
)

type UserProductRepI interface {
	AddRating(rating entities.Rating) error
	ChangeRating(rating entities.Rating) error
	GetRating(userid, productid string) error
	AddComment(c entities.Comment) error
	GetRatings(productid string) ([]entities.Rating, error)
	GetComments(productid string) ([]entities.Comment, error)
	AddToFaves(productid, userid string) error
	DeleteFromFaves(productid, userid string) error
	CheckFave(productid, userid string) error
	GetFaves(userid string) ([]entities.ProductLess, error)
	AddToCart(productid, userid string) error
	IsInCart(productid, userid string) (int, error)
	DeleteFromCart(productid, userid string, count int) error
	GetCart(userid string) ([]entities.Item, error)
	FinalizeCart(userid string) ([]entities.Item, error)
	AddInvoice(invoice entities.Invoice) error
	GetInvoices(userid string) ([]entities.Invoice, error)
}
