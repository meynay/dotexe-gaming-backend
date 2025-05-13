package user_rep

import (
	"fmt"
	"store/internal/entities"
	"store/pkg"
	"strings"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertUserByPhone(phone string) (*entities.User, error) {
	user := entities.User{
		Phone: phone,
		Faves: []uint{},
		Cart:  []entities.Item{},
	}
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return r.GetUserByPhone(phone)
}

func (r *UserRepository) InsertUserByEmail(email, password string) (*entities.User, error) {
	password, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}
	email = strings.ToLower(email)
	user := entities.User{
		Email:    email,
		Password: password,
		Faves:    []uint{},
		Cart:     []entities.Item{},
	}
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return r.GetUserByEmail(email)
}

func (r *UserRepository) GetUserByPhone(phone string) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, entities.User{Phone: phone})
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	email = strings.ToLower(email)
	err := r.db.First(&user, entities.User{Email: email})
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (r *UserRepository) CheckUser(email, password string) (*entities.User, error) {
	var user entities.User
	email = strings.ToLower(email)
	err := r.db.First(&user, entities.User{Email: email})
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	if pkg.CompareHashAndPassword(user.Password, password) != nil {
		return nil, fmt.Errorf("wrong password")
	}
	return &user, nil
}

func (r *UserRepository) SaveToken(userID uint, token string) error {
	res := r.db.Model(entities.User{}).Where("id = ?", userID).Update("refresh_roken = ?", token)
	return res.Error
}

func (r *UserRepository) TokenExists(userID uint, token string) error {
	var user entities.User
	tx := r.db.First(&user, userID)
	if tx.Error != nil {
		return tx.Error
	}
	if user.RefreshToken != token {
		return fmt.Errorf("wrong refresh token")
	}
	return nil
}
