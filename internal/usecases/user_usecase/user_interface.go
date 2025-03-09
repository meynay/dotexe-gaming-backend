package user_usecase

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserUsecaseI interface {
	FirstAttempt(inp string) (int, string)
	LoginWithEmail(email, password string) error
	LoginWithPhone(phone string) error
	SignupWithEmail(email, password string) error
	SignupWithPhone(Phone, code string) error
	SaveToken(id primitive.ObjectID, token string) error
	TokenExists(id primitive.ObjectID, token string) bool
}
