package cart_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/invoice_rep"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/user_rep"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartUsecase struct {
	productrep *product_rep.ProductRep
	userrep    *user_rep.UserRepository
	invoicerep *invoice_rep.InvoiceRep
}

func NewCartUsecase(pr *product_rep.ProductRep, ur *user_rep.UserRepository, ir *invoice_rep.InvoiceRep) *CartUsecase {
	return &CartUsecase{productrep: pr, userrep: ur, invoicerep: ir}
}

func (cu *CartUsecase) AddToCart(productid, userid primitive.ObjectID) error {
	return cu.userrep.AddToCart(productid, userid)
}

func (cu *CartUsecase) DeleteFromCart(productid, userid primitive.ObjectID) error {
	return cu.userrep.DeleteFromCart(productid, userid)
}

func (cu *CartUsecase) IncreaseInCart(productid, userid primitive.ObjectID) error {
	return cu.userrep.IncreaseInCart(productid, userid)
}

func (cu *CartUsecase) DecreaseInCart(productid, userid primitive.ObjectID) error {
	return cu.userrep.DecreaseInCart(productid, userid)
}

func (cu *CartUsecase) IsInCart(productid, userid primitive.ObjectID) (int, error) {
	return cu.userrep.IsInCart(productid, userid)
}

func (cu *CartUsecase) GetCart(userid primitive.ObjectID) []entities.CartItem {
	items, err := cu.userrep.GetCart(userid)
	if err != nil {
		return []entities.CartItem{}
	}
	var cartitems []entities.CartItem
	for _, item := range items {
		product, _ := cu.productrep.GetProduct(item.ProductID)
		cartitems = append(cartitems, entities.CartItem{
			ProductID:   item.ProductID,
			Count:       item.Count,
			ProductName: product.Name,
			Image:       product.Image,
			Price:       product.Price,
			Off:         product.Off,
		})
	}
	return cartitems
}

func (cu *CartUsecase) FinalizeCart(userid primitive.ObjectID, delivery_price int) error {
	items, err := cu.userrep.FinalizeCart(userid)
	if err != nil {
		return err
	}
	invoice := entities.Invoice{
		UserID:        userid,
		InvoiceDate:   time.Now(),
		OrderStatus:   entities.Processing,
		DeliveryPrice: delivery_price,
		TotalPrice:    1,
		Items:         items,
	}
	return cu.invoicerep.AddInvoice(invoice)
}
