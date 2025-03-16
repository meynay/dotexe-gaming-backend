package user_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecaseI interface {
	//user
	GetInfo(ID primitive.ObjectID) (entities.User, error)
	FillInfo(user entities.User) error
	ResetPassword(userid primitive.ObjectID, password string) error

	//login
	FirstAttempt(inp string) (int, string)
	LoginWithEmail(email, password string) error
	LoginWithPhone(phone string) error
	SignupWithEmail(email, password string) error
	SignupWithPhone(Phone, code string) error

	//refresh token
	SaveToken(id primitive.ObjectID, token string) error
	TokenExists(id primitive.ObjectID, token string) bool
}
