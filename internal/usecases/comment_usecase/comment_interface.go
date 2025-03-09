package comment_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUsecaseI interface {
	CommentOnProduct(c entities.Comment) error
	GetComments(productid primitive.ObjectID) []entities.CommentOut
}
