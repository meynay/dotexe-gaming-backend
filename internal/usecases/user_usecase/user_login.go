package user_usecase

import (
	"fmt"
	"log"
	"store/internal/repositories/user_rep"
	"store/pkg"
	"store/pkg/cacher"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	InvalidPhone = iota + 1
	InvalidEmail
	RegisterWithPhone
	RegisterWithEmail
	LoginWithPhone
	LoginWithEmail
)

type UserUsecase struct {
	userRep *user_rep.UserRepository
	cache   *cacher.Cacher
}

func NewUserUsecase(ur *user_rep.UserRepository, c *cacher.Cacher) *UserUsecase {
	return &UserUsecase{userRep: ur, cache: c}
}

func (u *UserUsecase) FirstAttempt(inp string) (int, string) {
	if pkg.IsNumeric(inp) {
		if !pkg.PhoneValidator(inp) {
			return InvalidPhone, "invalid phone"
		}
		user, err := u.userRep.GetUserByPhone(inp)
		number := pkg.RandomNumber()
		log.Println(number)
		u.cache.CacheSignInCode(inp, number)
		if err != nil {
			return RegisterWithPhone, "register with phone"
		}
		log.Println(user)
		return LoginWithPhone, "login with phone"
	}
	if !pkg.EmailValidator(inp) {
		return InvalidEmail, "invalid email"
	}
	user, err := u.userRep.GetUserByEmail(inp)
	if err != nil {
		return RegisterWithEmail, "register with email"
	}
	log.Println(user)
	return LoginWithEmail, "login with email"
}

func (u *UserUsecase) LoginWithEmail(email, password string) (primitive.ObjectID, error) {
	user, err := u.userRep.CheckUser(email, password)
	z, _ := primitive.ObjectIDFromHex("0")
	if err != nil {
		return z, err
	}
	return user.ID, nil
}

func (u *UserUsecase) LoginWithPhone(phone, code string) (primitive.ObjectID, error) {
	z, _ := primitive.ObjectIDFromHex("0")
	result, err := u.cache.CheckCode(phone, code)
	if err != nil {
		return z, err
	}
	if !result {
		return z, fmt.Errorf("wrong code")
	}
	user, err := u.userRep.GetUserByPhone(phone)
	if err != nil {
		return z, err
	}
	return user.ID, nil
}

func (u *UserUsecase) SignupWithEmail(email, password string) (primitive.ObjectID, error) {
	user, err := u.userRep.InsertUserByEmail(email, password)
	z, _ := primitive.ObjectIDFromHex("0")
	if err != nil {
		return z, err
	}
	return user.ID, nil
}

func (u *UserUsecase) SignupWithPhone(phone, code string) (primitive.ObjectID, error) {
	result, err := u.cache.CheckCode(phone, code)
	z, _ := primitive.ObjectIDFromHex("0")
	if err != nil {
		return z, err
	}
	if !result {
		return z, fmt.Errorf("wrong code")
	}
	user, err := u.userRep.InsertUserByPhone(phone)
	if err != nil {
		return z, err
	}
	return user.ID, nil
}

func (u *UserUsecase) SaveToken(id primitive.ObjectID, token string) error {
	return u.userRep.SaveToken(id, token)
}

func (u *UserUsecase) TokenExists(id primitive.ObjectID, token string) bool {
	return u.userRep.TokenExists(id, token) == nil
}
