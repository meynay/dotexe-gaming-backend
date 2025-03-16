package user_rep

import (
	"context"
	"fmt"
	"store/internal/entities"
	"store/pkg"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *UserRepository) GetPhoneNumber(userid primitive.ObjectID) string {
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": userid})
	var user entities.User
	res.Decode(&user)
	return user.Phone
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

func (ur *UserRepository) GetInfo(ID primitive.ObjectID) (entities.User, error) {
	res := ur.db.FindOne(context.TODO(), bson.M{"_id": ID})
	var user entities.User
	if err := res.Decode(&user); err != nil {
		return user, err
	}
	user.Password = ""
	return user, nil
}

func (ur *UserRepository) FillInfo(user entities.User) error {
	res := ur.db.FindOne(context.TODO(), bson.M{"email": strings.ToLower(user.Email)})
	var temp entities.User
	if res.Decode(&temp) == nil {
		if temp.ID != user.ID {
			return fmt.Errorf("email exists")
		}
	}
	res = ur.db.FindOne(context.TODO(), bson.M{"phone_number": user.Phone})
	if res.Decode(&temp) == nil {
		if temp.ID != user.ID {
			return fmt.Errorf("phone number exists")
		}
	}
	_, err := ur.db.UpdateOne(context.TODO(), bson.M{"_id": user.ID}, bson.M{
		"$set": bson.M{
			"phone_number": user.Phone,
			"email":        strings.ToLower(user.Email),
			"firstname":    user.FirstName,
			"lastname":     user.LastName,
			"address":      user.Address,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) ResetPassword(userid primitive.ObjectID, password string) error {
	pass, _ := pkg.HashPassword(password)
	_, err := ur.db.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$set": bson.M{"password": pass}})
	return err
}
