package blogpost_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
)

func (bp *BlogPostRep) AddComment(cm entities.BPComment) error {
	_, err := bp.bpcdb.InsertOne(context.TODO(), cm)
	if err != nil {
		return fmt.Errorf("couldn't add comment")
	}
	return nil
}

func (bp *BlogPostRep) GetComments(ID string) ([]entities.BPComment, error) {
	var comments []entities.BPComment
	cur, err := bp.bpcdb.Find(context.TODO(), bson.M{
		"blogpost_id": ID,
	})
	if err != nil {
		return comments, fmt.Errorf("couldn't find comments")
	}
	err = cur.Decode(&comments)
	if err != nil {
		return comments, fmt.Errorf("couldn't decode comments")
	}
	return comments, nil
}
