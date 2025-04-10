package category_rep

import (
	"context"
	"fmt"
	"log"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRep struct {
	rep *mongo.Collection
}

func NewCategoryRep(crep *mongo.Collection) *CategoryRep {
	return &CategoryRep{rep: crep}
}

func (cr *CategoryRep) AddCategory(c entities.Category) error {
	res, err := cr.rep.InsertOne(context.TODO(), c)
	if err != nil {
		return fmt.Errorf("couldn't insert category")
	}

	log.Printf("Category %s added with id %v\n", c.Name, res.InsertedID)
	return nil
}

func (cr *CategoryRep) EditCategory(c entities.Category) error {
	_, err := cr.rep.UpdateOne(context.TODO(), bson.M{"_id": c.ID}, bson.M{"$set": bson.M{
		"name":      c.Name,
		"image":     c.Image,
		"parent_id": c.ParentID,
	}})
	if err != nil {
		return fmt.Errorf("couldn't update category")
	}
	log.Printf("Category with id %s edited to name: %s \tand image: %s\n", c.ID, c.Name, c.Image)
	return nil
}

func (cr *CategoryRep) DeleteCategory(id primitive.ObjectID) error {
	_, err := cr.rep.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("couldn't delete category")
	}
	log.Printf("Category with id %s deleted\n", id)
	return nil
}

func (cr *CategoryRep) GetCategory(id primitive.ObjectID) (entities.Category, error) {
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
	result, err := cr.rep.Find(context.TODO(), bson.D{})
	if err != nil {
		return categories
	}
	result.All(context.TODO(), &categories)
	return categories
}

func (cr *CategoryRep) GetParents(ID primitive.ObjectID) []primitive.ObjectID {
	parents := []primitive.ObjectID{}
	res := cr.rep.FindOne(context.TODO(), bson.M{"_id": ID})
	var c entities.Category
	res.Decode(&c)
	z, _ := primitive.ObjectIDFromHex("0")
	for c.ParentID != z {
		parents = append(parents, c.ParentID)
		res = cr.rep.FindOne(context.TODO(), bson.M{"_id": c.ParentID})
		res.Decode(&c)
	}
	return parents
}
