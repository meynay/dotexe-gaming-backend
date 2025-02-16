package user_rep

import (
	"context"
	"fmt"
	"store/internal/entities"
	"store/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) InsertUserByPhone(phone string) (*entities.User, error) {
	user := bson.M{
		"phone":      phone,
		"created_at": time.Now(),
	}
	_, err := r.db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return r.GetUserByPhone(phone)
}

func (r *UserRepository) InsertUserByEmail(email, password string) (*entities.User, error) {
	password, err := pkg.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := bson.M{
		"email":      email,
		"password":   password,
		"created_at": time.Now(),
	}
	_, err = r.db.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return r.GetUserByEmail(email)
}

func (r *UserRepository) GetUserByPhone(phone string) (*entities.User, error) {
	var user entities.User
	err := r.db.FindOne(context.TODO(), bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (r *UserRepository) CheckUser(email, password string) (*entities.User, error) {
	var user entities.User
	err := r.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	if pkg.CompareHashAndPassword(user.Password, password) != nil {
		return nil, fmt.Errorf("wrong password")
	}
	return &user, nil
}

func (r *UserRepository) SaveToken(userID, token string) error {
	_, err := r.db.UpdateOne(
		context.TODO(),
		bson.M{"id": userID},
		bson.M{"refresh_token": token},
	)
	return err
}

func (r *UserRepository) TokenExists(userID, token string) error {
	var user entities.User
	err := r.db.FindOne(context.TODO(), bson.M{"id": userID, "refresh_token": token}).Decode(&user)
	return err
}
