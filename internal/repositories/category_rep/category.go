package category_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRep struct {
	rep *mongo.Collection
}

func NewCategoryRep(crep *mongo.Collection) *CategoryRep {
	return &CategoryRep{rep: crep}
}

func (cr *CategoryRep) AddCategory(c entities.Category) error {
	_, err := cr.rep.InsertOne(context.TODO(), c)
	if err != nil {
		return fmt.Errorf("couldn't insert category")
	}
	return nil
}

func (cr *CategoryRep) EditCategory(c entities.Category) error {
	_, err := cr.rep.UpdateOne(context.TODO(), bson.M{"_id": c.ID}, c)
	if err != nil {
		return fmt.Errorf("couldn't update category")
	}
	return nil
}

func (cr *CategoryRep) DeleteCategory(id string) error {
	_, err := cr.rep.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("couldn't delete category")
	}
	return nil
}

func (cr *CategoryRep) GetCategory(id string) (entities.Category, error) {
	result := cr.rep.FindOne(context.TODO(), bson.M{"_id": id})
	if result.Err() != nil {
		return entities.Category{}, fmt.Errorf("couldn't find category")
	}
	c := entities.Category{}
	result.Decode(&c)
	return c, nil
}

func (cr *CategoryRep) GetCategories() []entities.Category {
	categories := []entities.Category{}
	result, err := cr.rep.Find(context.TODO(), bson.M{})
	if err != nil {
		return categories
	}
	result.Decode(&categories)
	return categories
}

func (cr *CategoryRep) GetParents(ID string) []string {
	parents := []string{}
	res := cr.rep.FindOne(context.TODO(), bson.M{"_id": ID})
	var c entities.Category
	res.Decode(&c)
	for c.ParentID != "" {
		parents = append(parents, c.ParentID)
		res = cr.rep.FindOne(context.TODO(), bson.M{"_id": c.ParentID})
		res.Decode(&c)
	}
	return parents
}
