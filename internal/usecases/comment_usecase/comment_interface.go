package comment_usecase

import "store/internal/entities"

type CommentUsecaseI interface {
	CommentOnProduct(c entities.Comment) error
	GetComments(productid string) []entities.CommentOut
}
