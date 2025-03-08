package user_rep

import "store/internal/entities"

type UserRepositoryI interface {
	//login-signup
	InsertUserByPhone(phone string) error
	InsertUserByEmail(email, password string) error
	GetUserByPhone(phone string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckUser(email, password string) (*entities.User, error)
	SaveToken(userID, token string) error

	//faves
	AddToFaves(productid, userid string) error
	DeleteFromFaves(productid, userid string) error
	CheckFave(productid, userid string) error
	GetFaves(userid string) []string

	//cart
	AddToCart(productid, userid string) error
	IsInCart(productid, userid string) (int, error)
	DeleteFromCart(productid, userid string, count int) error
	GetCart(userid string) []entities.Item
	FinalizeCart(userid string) []entities.Item
}
