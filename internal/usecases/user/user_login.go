package user_usecase

import (
	"log"
	user_rep "store/internal/repositories/user"
	"store/pkg"
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
}

func NewUserUsecase(ur *user_rep.UserRepository) *UserUsecase {
	return &UserUsecase{userRep: ur}
}

func (u *UserUsecase) FirstAttempt(inp string) (int, string) {
	if pkg.IsNumeric(inp) {
		if !pkg.PhoneValidator(inp) {
			return InvalidPhone, "invalid phone"
		}
		user, err := u.userRep.GetUserByPhone(inp)
		number := pkg.RandomNumber()
		log.Println(number)
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

func (u *UserUsecase) LoginWithEmail(email, password string) (string, error) {
	user, err := u.userRep.CheckUser(email, password)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *UserUsecase) LoginWithPhone(phone, code string) (string, error) {
	user, err := u.userRep.GetUserByPhone(phone)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *UserUsecase) SignupWithEmail(email, password string) (string, error) {
	user, err := u.userRep.InsertUserByEmail(email, password)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *UserUsecase) SignupWithPhone(Phone, code string) (string, error) {
	user, err := u.userRep.InsertUserByPhone(Phone)
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *UserUsecase) SaveToken(id, token string) error {
	return u.userRep.SaveToken(id, token)
}

func (u *UserUsecase) TokenExists(id, token string) bool {
	return u.userRep.TokenExists(id, token) == nil
}
