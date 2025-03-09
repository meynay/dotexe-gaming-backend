package comment_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepI interface {
	AddComment(c entities.Comment) error
	GetComments(productid primitive.ObjectID) ([]entities.Comment, error)
}
