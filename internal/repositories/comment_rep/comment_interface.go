package comment_rep

import "store/internal/entities"

type CommentRepI interface {
	AddComment(c entities.Comment) error
	GetComments(productid string) ([]entities.Comment, error)
}
