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
	res := pr.cgdb.FindOne(context.TODO(), bson.M{"_id": ID})
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
	res := pr.cgdb.FindOne(context.TODO(), bson.M{"_id": ID})
	var c entities.Category
	res.Decode(&c)
	for c.ParentID != "" {
		parents = append(parents, c.ParentID)
		res = pr.cgdb.FindOne(context.TODO(), bson.M{"_id": c.ParentID})
		res.Decode(&c)
	}
	return parents
}

func (pr *ProductRep) EditCategory(c entities.Category) error {
	_, err := pr.prdb.UpdateOne(context.TODO(), bson.M{
		"_id": c.ID,
	}, c)
	if err != nil {
		return fmt.Errorf("couldn't update category")
	}
	return nil
}

func (pr *ProductRep) DeleteCategory(ID string) error {
	_, err := pr.cgdb.DeleteOne(context.TODO(), bson.M{"_id": ID})
	if err != nil {
		return fmt.Errorf("couldn't delete category")
	}
	return nil
}

func (pr *ProductRep) GetCategories() []entities.Category {
	categories := []entities.Category{}
	cur, err := pr.cgdb.Find(context.TODO(), bson.M{})
	if err != nil {
		return categories
	}
	cur.Decode(&categories)
	return categories
}
