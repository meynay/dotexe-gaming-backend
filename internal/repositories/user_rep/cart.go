package user_rep

import (
	"context"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
)

func (ur *UserRepository) AddToCart(productid, userid string) error {
	item := entities.Item{
		ProductID: productid,
		Count:     1,
	}
	ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$push": bson.M{"cart": item}})
	return nil
}

func (ur *UserRepository) EditQuantity(productid, userid string, asc bool) error {
	var update bson.M
	if asc {
		update = bson.M{}
	} else {
		update = bson.M{}
	}

	ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid, "cart.product_id": productid}, update)
	return nil
}

func (ur *UserRepository) IsInCart(productid, userid string) (int, error) {
	return 0, nil
}

func (ur *UserRepository) GetCart(userid string) ([]entities.Item, error) {
	user := entities.User{}
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	res.Decode(&user)
	return user.Cart, nil
}

func (ur *UserRepository) FinalizeCart(userid string) ([]entities.Item, error) {
	user := entities.User{}
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	res.Decode(&user)
	ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$set": bson.M{"cart": []entities.Item{}}})
	return user.Cart, nil
}
