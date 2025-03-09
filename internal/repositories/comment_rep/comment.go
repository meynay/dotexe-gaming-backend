package comment_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentRep struct {
	rep *mongo.Collection
}

func NewCommentRep(crep *mongo.Collection) *CommentRep {
	return &CommentRep{rep: crep}
}

func (cr *CommentRep) AddComment(c entities.Comment) error {
	_, err := cr.rep.InsertOne(context.TODO(), c)
	if err != nil {
		return fmt.Errorf("couldn't insert comment")
	}
	return nil
}

func (cr *CommentRep) GetComments(productid primitive.ObjectID) ([]entities.Comment, error) {
	var comments []entities.Comment
	res, err := cr.rep.Find(context.TODO(), bson.M{"product_id": productid})
	if err != nil {
		return comments, fmt.Errorf("couldn't find comments on product")
	}
	res.All(context.TODO(), &comments)
	return comments, nil
}
