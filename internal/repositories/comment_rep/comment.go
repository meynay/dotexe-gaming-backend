package comment_rep

import (
	"fmt"
	"store/internal/entities"

	"gorm.io/gorm"
)

type CommentRep struct {
	rep *gorm.DB
}

func NewCommentRep(crep *gorm.DB) *CommentRep {
	return &CommentRep{rep: crep}
}

func (cr *CommentRep) AddComment(c entities.Comment) error {
	tx := cr.rep.Create(&c)
	if tx.Error != nil {
		return fmt.Errorf("couldn't insert comment")
	}
	return nil
}

func (cr *CommentRep) GetComments(productid uint) ([]entities.Comment, error) {
	var comments []entities.Comment
	tx := cr.rep.First(&comments, entities.Comment{ProductID: productid})
	if tx.Error != nil {
		return comments, fmt.Errorf("couldn't find comments on product")
	}
	return comments, nil
}
