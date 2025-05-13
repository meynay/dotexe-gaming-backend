package blogpost_rep

import (
	"fmt"
	"store/internal/entities"
)

func (bp *BlogPostRep) AddComment(cm entities.BPComment) error {
	tx := bp.bpdb.Create(&cm)
	if tx.Error != nil {
		return fmt.Errorf("couldn't add comment")
	}
	return nil
}

func (bp *BlogPostRep) GetComments(ID uint) ([]entities.BPComment, error) {
	var comments []entities.BPComment
	tx := bp.bpdb.Find(&comments)
	if tx.Error != nil {
		return comments, fmt.Errorf("couldn't find comments")
	}
	out := []entities.BPComment{}
	for _, comment := range comments {
		if comment.BlogPostID == ID {
			out = append(out, comment)
		}
	}
	return out, nil
}
