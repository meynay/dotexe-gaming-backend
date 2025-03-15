package cart_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartUsecaseI interface {
	//cart actions
	AddToCart(productid, userid primitive.ObjectID) error
	DeleteFromCart(productid, userid primitive.ObjectID) error
	IncreaseInCart(productid, userid primitive.ObjectID) error
	DecreaseInCart(productid, userid primitive.ObjectID) error
	IsInCart(productid, userid primitive.ObjectID) (int, error)
	GetCart(userid primitive.ObjectID) []entities.CartItem
	FinalizeCart(userid primitive.ObjectID) error

	//order actions
	GetInvoices(userid primitive.ObjectID) []entities.Invoice
	GetInvoice(userid, invoiceid primitive.ObjectID) entities.Invoice
}
