package fave_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FaveUsecaseI interface {
	FaveProduct(userid, productid primitive.ObjectID) error
	UnfaveProduct(userid, productid primitive.ObjectID) error
	CheckFave(userid, productid primitive.ObjectID) error
	GetFaves(userid primitive.ObjectID) []entities.ProductLess
}
