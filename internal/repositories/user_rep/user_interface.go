package user_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepositoryI interface {
	//login-signup
	InsertUserByPhone(phone string) error
	InsertUserByEmail(email, password string) error
	GetUserByPhone(phone string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckUser(email, password string) (*entities.User, error)
	SaveToken(userID primitive.ObjectID, token string) error

	//faves
	AddToFaves(productid, userid primitive.ObjectID) error
	DeleteFromFaves(productid, userid primitive.ObjectID) error
	CheckFave(productid, userid primitive.ObjectID) error
	GetFaves(userid primitive.ObjectID) []primitive.ObjectID

	//cart
	AddToCart(productid, userid primitive.ObjectID) error
	IsInCart(productid, userid primitive.ObjectID) (int, error)
	DeleteFromCart(productid, userid primitive.ObjectID, count int) error
	GetCart(userid primitive.ObjectID) []entities.Item
	FinalizeCart(userid primitive.ObjectID) []entities.Item
}
