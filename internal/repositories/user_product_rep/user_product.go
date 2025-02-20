package user_product_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserProductRep struct {
	rates    *mongo.Collection
	user     *mongo.Collection
	product  *mongo.Collection
	invoices *mongo.Collection
	comments *mongo.Collection
	category *mongo.Collection
}

func NewUserProductRep(r, u, p, i, c, c2 *mongo.Collection) *UserProductRep {
	return &UserProductRep{rates: r, user: u, product: p, invoices: i, comments: c, category: c2}
}

func (up *UserProductRep) AddRating(rating entities.Rating) error {
	_, err := up.rates.InsertOne(context.TODO(), rating)
	if err != nil {
		return fmt.Errorf("couldn't insert rate")
	}
	return nil
}

func (up *UserProductRep) ChangeRating(rating entities.Rating) error {
	_, err := up.rates.DeleteOne(context.TODO(), bson.M{"user_id": rating.UserID, "product_id": rating.ProductID})
	if err != nil {
		return fmt.Errorf("can't delete existing rate")
	}
	_, err = up.rates.InsertOne(context.TODO(), rating)
	if err != nil {
		return fmt.Errorf("couldn't insert rate")
	}
	return nil
}

func (up *UserProductRep) GetRating(userid, productid string) (entities.Rating, error) {
	res := up.rates.FindOne(context.TODO(), bson.M{"user_id": userid, "product_id": productid})
	if res.Err() != nil {
		return entities.Rating{}, res.Err()
	}
	var r entities.Rating
	res.Decode(&r)
	return r, nil
}

func (up *UserProductRep) AddComment(c entities.Comment) error {
	_, err := up.comments.InsertOne(context.TODO(), c)
	if err != nil {
		return fmt.Errorf("couldn't insert comment")
	}
	return nil
}

func (up *UserProductRep) GetRatings(productid string) ([]entities.Rating, error) {
	var rates []entities.Rating
	res, err := up.rates.Find(context.TODO(), bson.M{"product_id": productid})
	if err != nil {
		return rates, fmt.Errorf("couldn't find rates on product")
	}
	res.Decode(&rates)
	return rates, nil
}

func (up *UserProductRep) GetComments(productid string) ([]entities.Comment, error) {
	var comments []entities.Comment
	res, err := up.comments.Find(context.TODO(), bson.M{"product_id": productid})
	if err != nil {
		return comments, fmt.Errorf("couldn't find comments on product")
	}
	res.Decode(&comments)
	return comments, nil
}

func (up *UserProductRep) AddToFaves(productid, userid string) error {
	_, err := up.user.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$addToSet": bson.M{"faves": productid}})
	if err != nil {
		return fmt.Errorf("couldn't add product to faves")
	}
	return nil
}

func (up *UserProductRep) DeleteFromFaves(productid, userid string) error {
	_, err := up.user.UpdateOne(context.TODO(), bson.M{"_id": userid}, bson.M{"$pull": bson.M{"faves": productid}})
	if err != nil {
		return fmt.Errorf("couldn't delete product from faves")
	}
	return nil
}

func (up *UserProductRep) CheckFave(productid, userid string) error {
	res := up.user.FindOne(context.TODO(), bson.M{"_id": userid, "faves.product_id": productid})
	if res.Err() != nil {
		return fmt.Errorf("couldn't find product at user faves")
	}
	return nil
}

func (up *UserProductRep) GetFaves(userid string) ([]entities.Product, error) {
	var products []entities.Product
	res := up.user.FindOne(context.TODO(), bson.M{"_id": userid})
	var u entities.User
	res.Decode(&u)
	result, err := up.product.Find(context.TODO(), bson.M{"_id": bson.M{"$in": u.Faves}})
	if err != nil {
		return products, fmt.Errorf("couldn't find products")
	}
	result.Decode(&products)
	return products, nil
}

func (up *UserProductRep) GetCategoryName(ID string) string {
	res := up.category.FindOne(context.TODO(), bson.M{"_id": ID})
	if res.Err() != nil {
		return ""
	}
	var c entities.Category
	if res.Decode(&c) != nil {
		return ""
	}
	return c.Name
}

func (up *UserProductRep) GetUsername(userid string) string {
	res := up.user.FindOne(context.TODO(), bson.M{"_id": userid})
	var user entities.User
	res.Decode(&user)
	if user.FirstName != "" {
		return user.FirstName + " " + user.LastName
	}
	return "ناشناس"
}
