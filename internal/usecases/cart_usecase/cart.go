package cart_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/invoice_rep"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/user_rep"
)

type CartUsecase struct {
	productrep *product_rep.ProductRep
	userrep    *user_rep.UserRepository
	invoicerep *invoice_rep.InvoiceRep
}

func NewCartUsecase(pr *product_rep.ProductRep, ur *user_rep.UserRepository, ir *invoice_rep.InvoiceRep) *CartUsecase {
	return &CartUsecase{productrep: pr, userrep: ur, invoicerep: ir}
}

func (cu *CartUsecase) AddToCart(productid, userid uint) error {
	return cu.userrep.AddToCart(productid, userid)
}

func (cu *CartUsecase) DeleteFromCart(productid, userid uint) error {
	return cu.userrep.DeleteFromCart(productid, userid)
}

func (cu *CartUsecase) IncreaseInCart(productid, userid uint) error {
	return cu.userrep.IncreaseInCart(productid, userid)
}

func (cu *CartUsecase) DecreaseInCart(productid, userid uint) error {
	return cu.userrep.DecreaseInCart(productid, userid)
}

func (cu *CartUsecase) IsInCart(productid, userid uint) (int, error) {
	return cu.userrep.IsInCart(productid, userid)
}

func (cu *CartUsecase) GetCart(userid uint) []entities.CartItem {
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

func (cu *CartUsecase) FinalizeCart(userid uint, delivery_price int) error {
	items, err := cu.userrep.FinalizeCart(userid)
	if err != nil {
		return err
	}
	its := []entities.Item{}
	for _, val := range items {
		its = append(its, entities.Item{
			ProductID: val.ProductID,
			Count:     val.Count,
			Price:     val.Price,
			Off:       val.Off,
		})
	}
	invoice := entities.Invoice{
		UserID:        userid,
		OrderStatus:   entities.Processing,
		DeliveryPrice: delivery_price,
		TotalPrice:    1,
		Items:         its,
	}
	return cu.invoicerep.AddInvoice(invoice)
}
