package product_rep

import (
	"context"
	"fmt"
	"store/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRep struct {
	prdb *mongo.Collection
}

func NewProductRep(pr *mongo.Collection) *ProductRep {
	return &ProductRep{prdb: pr}
}

func (pr *ProductRep) AddProduct(p entities.Product) error {
	_, err := pr.prdb.InsertOne(context.TODO(), p)
	if err != nil {
		return fmt.Errorf("couldn't add product")
	}
	return nil
}
func (pr *ProductRep) GetProduct(ID primitive.ObjectID) (entities.Product, error) {
	res := pr.prdb.FindOne(context.TODO(), bson.M{
		"_id": ID,
	})
	var product entities.Product
	if res.Err() != nil {
		return product, fmt.Errorf("couldn't find product")
	}
	if res.Decode(&product) != nil {
		return product, fmt.Errorf("couldn't decode product")
	}
	return product, nil
}
func (pr *ProductRep) GetProducts() ([]entities.Product, error) {
	cur, err := pr.prdb.Find(context.TODO(), bson.M{})
	var products []entities.Product
	if err != nil {
		return products, fmt.Errorf("couldn't get products")
	}
	if cur.All(context.TODO(), &products) != nil {
		return products, fmt.Errorf("couldn't decode products")
	}
	return products, nil
}

func (pr *ProductRep) EditProduct(p entities.Product) error {
	_, err := pr.prdb.UpdateOne(context.TODO(), bson.M{
		"_id": p.ID,
	}, bson.M{"$set": bson.M{
		"name":        p.Name,
		"image":       p.Image,
		"images":      p.Images,
		"description": p.Description,
		"price":       p.Price,
		"stock":       p.Stock,
		"info":        p.Info,
		"off":         p.Off,
		"category_id": p.CategoryID,
		"tags":        p.Tags,
		"updated_at":  time.Now(),
	}})
	if err != nil {
		return fmt.Errorf("couldn't update product")
	}
	return nil
}

func (pr *ProductRep) DeleteProduct(ID primitive.ObjectID) error {
	_, err := pr.prdb.DeleteOne(context.TODO(), bson.M{"_id": ID})
	if err != nil {
		return fmt.Errorf("couldn't delete product")
	}
	return nil
}
