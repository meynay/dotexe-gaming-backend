package product_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
)

func (pr *ProductRep) AddCategory(c entities.Category) error {
	_, err := pr.cgdb.InsertOne(context.TODO(), c)
	if err != nil {
		return fmt.Errorf("couldn't add category")
	}
	return nil
}

func (pr *ProductRep) GetCategoryName(ID string) (string, error) {
	res := pr.cgdb.FindOne(context.TODO(), bson.M{"id": ID})
	if res.Err() != nil {
		return "", fmt.Errorf("couldn't find category")
	}
	var c entities.Category
	if res.Decode(&c) != nil {
		return "", fmt.Errorf("couldn't decode category")
	}
	return c.Name, nil
}

func (pr *ProductRep) GetParents(ID string) []string {
	parents := []string{}
	res := pr.cgdb.FindOne(context.TODO(), bson.M{"id": ID})
	var c entities.Category
	res.Decode(&c)
	for c.ParentID != "" {
		parents = append(parents, c.ParentID)
		res = pr.cgdb.FindOne(context.TODO(), bson.M{"id": c.ParentID})
		res.Decode(&c)
	}
	return parents
}
