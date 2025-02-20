package blogpost_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogPostRep struct {
	bpdb  *mongo.Collection
	bpcdb *mongo.Collection
}

func NewBlogPostRep(bp *mongo.Collection, bpc *mongo.Collection) *BlogPostRep {
	return &BlogPostRep{bpdb: bp, bpcdb: bpc}
}

func (bp *BlogPostRep) AddBlogPost(b entities.BlogPost) error {
	_, err := bp.bpdb.InsertOne(context.TODO(), b)
	if err != nil {
		return fmt.Errorf("couldn't insert blogpost")
	}
	return nil
}

func (bp *BlogPostRep) GetBlogPost(ID string) (entities.BlogPost, error) {
	res := bp.bpdb.FindOne(context.TODO(), bson.M{
		"_id": ID,
	})
	if res.Err() != nil {
		return entities.BlogPost{}, fmt.Errorf("no documents found")
	}
	var blogpost entities.BlogPost
	err := res.Decode(&blogpost)
	if err != nil {
		return entities.BlogPost{}, fmt.Errorf("unable to decode results")
	}
	return blogpost, nil
}

func (bp *BlogPostRep) GetBlogPosts() ([]entities.BlogPost, error) {
	cur, err := bp.bpdb.Find(context.TODO(), bson.M{})
	var posts []entities.BlogPost
	if err != nil {
		return posts, err
	}
	err = cur.Decode(&posts)
	if err != nil {
		return posts, fmt.Errorf("can't decode blogposts")
	}
	return posts, nil
}

func (bp *BlogPostRep) EditBlogPost(b entities.BlogPost) error {
	_, err := bp.bpdb.UpdateOne(context.TODO(), bson.M{
		"_id": b.ID,
	}, b)
	if err != nil {
		return fmt.Errorf("couldn't update blogpost")
	}
	return nil
}

func (bp *BlogPostRep) DeleteBlogPost(ID string) error {
	_, err := bp.bpdb.DeleteOne(context.TODO(), bson.M{"_id": ID})
	if err != nil {
		return fmt.Errorf("couldn't delete blogpost")
	}
	return nil
}
