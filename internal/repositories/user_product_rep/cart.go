package user_product_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
)

func (up *UserProductRep) AddToCart(productid, userid string) error {
	item := entities.Item{
		ProductID: productid,
		Count:     1,
	}
	up.user.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$push": bson.M{"cart": item}})
	return nil
}

func (up *UserProductRep) EditQuantity(productid, userid string, asc bool) error {
	var update bson.M
	if asc {
		update = bson.M{}
	} else {
		update = bson.M{}
	}

	up.user.UpdateOne(context.TODO(), bson.M{"_id": userid, "cart.product_id": productid}, update)
	return nil
}

func (up *UserProductRep) IsInCart(productid, userid string) (int, error) {
	return 0, nil
}

func (up *UserProductRep) GetCart(userid string) ([]entities.Item, error) {
	user := entities.User{}
	res := up.user.FindOne(context.TODO(), bson.M{"_id": userid})
	res.Decode(&user)
	return user.Cart, nil
}

func (up *UserProductRep) FinalizeCart(userid string) ([]entities.Item, error) {
	user := entities.User{}
	res := up.user.FindOne(context.TODO(), bson.M{"_id": userid})
	res.Decode(&user)
	up.user.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$set": bson.M{"cart": []entities.Item{}}})
	return user.Cart, nil
}

func (up *UserProductRep) AddInvoice(invoice entities.Invoice) error {
	_, err := up.invoices.InsertOne(context.TODO(), invoice)
	if err != nil {
		return fmt.Errorf("couldn't insert invoice")
	}
	return nil
}

func (up *UserProductRep) GetInvoices(userid string) ([]entities.Invoice, error) {
	invoices := []entities.Invoice{}
	res, err := up.invoices.Find(context.TODO(), bson.M{"user_id": userid})
	if err != nil {
		return invoices, fmt.Errorf("error getting invoices")
	}
	res.Decode(&invoices)
	return invoices, nil
}
