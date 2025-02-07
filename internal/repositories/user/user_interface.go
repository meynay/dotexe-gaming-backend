package user_rep

import "store/internal/entities"

type UserRepositoryI interface {
	InsertUserByPhone(phone string) error
	InsertUserByEmail(email, password string) error
	GetUserByPhone(phone string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckUser(email, password string) (*entities.User, error)
	SaveToken(userID, token string) error
}
