package comment_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/comment_rep"
	"store/internal/repositories/user_rep"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentUsecase struct {
	commentrep *comment_rep.CommentRep
	userrep    *user_rep.UserRepository
}

func NewCommentUsecase(cr *comment_rep.CommentRep, ur *user_rep.UserRepository) *CommentUsecase {
	return &CommentUsecase{commentrep: cr, userrep: ur}
}

func (cu *CommentUsecase) CommentOnProduct(c entities.Comment) error {
	return cu.commentrep.AddComment(c)
}

func (cu *CommentUsecase) GetComments(productid primitive.ObjectID) []entities.CommentOut {
	cmnt, err := cu.commentrep.GetComments(productid)
	comments := []entities.CommentOut{}
	if err != nil {
		return comments
	}
	for _, c := range cmnt {
		newcomment := entities.CommentOut{
			ID:        c.ID,
			Parent:    c.Parent,
			User:      cu.userrep.GetUsername(c.UserID),
			Comment:   c.Comment,
			CreatedAt: c.CreatedAt,
		}
		comments = append(comments, newcomment)
	}
	return comments
}
