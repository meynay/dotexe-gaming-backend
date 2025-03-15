package user_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (ur *UserRepository) AddToCart(productid, userid primitive.ObjectID) error {
	item := entities.Item{
		ProductID: productid,
		Count:     1,
	}
	ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$push": bson.M{"cart": item}})
	return nil
}

func (ur *UserRepository) DeleteFromCart(productid, userid primitive.ObjectID) error {
	_, err := ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$pull": bson.M{"cart": bson.M{"product_id": productid}}})
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) IsInCart(productid, userid primitive.ObjectID) (int, error) {
	return 0, nil
}

func (ur *UserRepository) IncreaseInCart(productid, userid primitive.ObjectID) error {
	filter := bson.M{"_id": userid}
	update := bson.M{
		"$inc": bson.M{"cart.$[elem].count": 1},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.product_id": productid}},
	}
	opts := options.Update().SetArrayFilters(arrayFilters)
	_, err := ur.db.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to update cart: %w", err)
	}
	return nil
}

func (ur *UserRepository) DecreaseInCart(productid, userid primitive.ObjectID) error {
	filter := bson.M{"_id": userid}
	update := bson.M{
		"$inc": bson.M{"cart.$[elem].count": -1},
	}
	arrayFilters := options.ArrayFilters{
		Filters: []interface{}{bson.M{"elem.product_id": productid}},
	}
	opts := options.Update().SetArrayFilters(arrayFilters)
	_, err := ur.db.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to update cart: %w", err)
	}
	return nil
}

func (ur *UserRepository) GetCart(userid primitive.ObjectID) ([]entities.Item, error) {
	user := entities.User{}
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	res.Decode(&user)
	return user.Cart, nil
}

func (ur *UserRepository) FinalizeCart(userid primitive.ObjectID) ([]entities.Item, error) {
	user := entities.User{}
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	res.Decode(&user)
	ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$set": bson.M{"cart": []entities.Item{}}})
	return user.Cart, nil
}
