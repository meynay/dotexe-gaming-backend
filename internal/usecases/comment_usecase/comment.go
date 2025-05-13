package comment_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/admin_rep"
	"store/internal/repositories/comment_rep"
	"store/internal/repositories/user_rep"
)

type CommentUsecase struct {
	commentrep *comment_rep.CommentRep
	userrep    *user_rep.UserRepository
	adminrep   *admin_rep.AdminRep
}

func NewCommentUsecase(cr *comment_rep.CommentRep, ur *user_rep.UserRepository, ar *admin_rep.AdminRep) *CommentUsecase {
	return &CommentUsecase{commentrep: cr, userrep: ur, adminrep: ar}
}

func (cu *CommentUsecase) CommentOnProduct(c entities.Comment) error {
	return cu.commentrep.AddComment(c)
}

func (cu *CommentUsecase) GetComments(productid uint) []entities.CommentOut {
	cmnt, err := cu.commentrep.GetComments(productid)
	comments := []entities.CommentOut{}
	if err != nil {
		return comments
	}
	for _, c := range cmnt {
		var newcomment entities.CommentOut
		if c.IsAdmin {
			newcomment = entities.CommentOut{
				ID:        c.ID,
				CreatedAt: c.CreatedAt,
				Parent:    *c.ParentID,
				User:      cu.adminrep.GetName(c.UserID),
				Comment:   c.Comment,
			}
		} else {
			newcomment = entities.CommentOut{
				Parent:    *c.ParentID,
				ID:        c.ID,
				CreatedAt: c.CreatedAt,
				User:      cu.userrep.GetUsername(c.UserID),
				Comment:   c.Comment,
			}
		}
		comments = append(comments, newcomment)
	}
	return comments
}
