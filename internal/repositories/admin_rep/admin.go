package admin_rep

import (
	"context"
	"fmt"
	"store/internal/entities"
	"store/pkg"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRep struct {
	rep *mongo.Collection
}

func NewAdminRep(ar *mongo.Collection) *AdminRep {
	return &AdminRep{rep: ar}
}

func (ar *AdminRep) GetName(id primitive.ObjectID) string {
	res := ar.rep.FindOne(context.TODO(), bson.M{"_id": id})
	admin := entities.Admin{}
	if err := res.Decode(&admin); err != nil {
		return err.Error()
	}
	if admin.FirstName == "" && admin.LastName == "" {
		return "ادمین"
	}
	return admin.FirstName + " " + admin.LastName + " (ادمین)"
}

func (ar *AdminRep) AddAdmin(username, password string) error {
	res := ar.rep.FindOne(context.TODO(), bson.M{"username": username})
	admin := entities.Admin{}
	res.Decode(&admin)
	username = strings.ToLower(username)
	if admin.Username == username {
		return fmt.Errorf("this username exists")
	}
	pass, err := pkg.HashPassword(password)
	if err != nil {
		return err
	}
	admin = entities.Admin{
		Password: pass,
		Username: username,
	}
	_, err = ar.rep.InsertOne(context.TODO(), admin)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AdminRep) GetInfo(adminID primitive.ObjectID) (entities.Admin, error) {
	res := ar.rep.FindOne(context.TODO(), bson.M{"_id": adminID})
	admin := entities.Admin{}
	err := res.Decode(&admin)
	return admin, err
}

func (ar *AdminRep) FillFields(admin entities.Admin) error {
	res := ar.rep.FindOne(context.TODO(), bson.M{"phone_number": admin.Phone})
	var temp entities.Admin
	if res.Decode(&temp) == nil {
		if temp.ID != admin.ID {
			return fmt.Errorf("phone number exists")
		}
	}
	_, err := ar.rep.UpdateOne(context.TODO(), bson.M{"_id": admin.ID},
		bson.M{
			"$set": bson.M{
				"phone_number": admin.Phone,
				"firstname":    admin.FirstName,
				"lastname":     admin.LastName,
				"image":        admin.Image,
				"bio":          admin.Bio,
			}},
	)
	if err != nil {
		return err
	}
	return nil
}

func (ar *AdminRep) Login(username, password string) (primitive.ObjectID, error) {
	username = strings.ToLower(username)
	res := ar.rep.FindOne(context.TODO(), bson.M{"username": username})
	admin := entities.Admin{}
	if err := res.Decode(&admin); err != nil {
		id, _ := primitive.ObjectIDFromHex("0")
		return id, err
	}
	if pkg.CompareHashAndPassword(admin.Password, password) != nil {
		id, _ := primitive.ObjectIDFromHex("1")
		return id, fmt.Errorf("")
	}
	return admin.ID, nil
}
