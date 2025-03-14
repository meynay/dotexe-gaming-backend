package user_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) AddToFaves(productid, userid primitive.ObjectID) error {
	_, err := ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$addToSet": bson.M{"faves": productid}})
	if err != nil {
		return fmt.Errorf("couldn't add product to faves")
	}
	return nil
}

func (ur *UserRepository) DeleteFromFaves(productid, userid primitive.ObjectID) error {
	_, err := ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$pull": bson.M{"faves": productid}})
	if err != nil {
		return fmt.Errorf("couldn't delete product from faves")
	}
	return nil
}

func (ur *UserRepository) CheckFave(productid, userid primitive.ObjectID) error {
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid, "faves.product_id": productid})
	if res.Err() != nil {
		return fmt.Errorf("couldn't find product at user faves")
	}
	return nil
}

func (ur *UserRepository) GetFaves(userid primitive.ObjectID) []primitive.ObjectID {
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	var u entities.User
	res.Decode(&u)
	return u.Faves
}

func (ur *UserRepository) GetUsername(userid primitive.ObjectID) string {
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	var user entities.User
	res.Decode(&user)
	if user.FirstName != "" {
		return user.FirstName + " " + user.LastName
	}
	return "ناشناس"
}
