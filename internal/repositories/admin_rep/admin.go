package admin_rep

import "go.mongodb.org/mongo-driver/mongo"

type AdminRep struct {
	rep *mongo.Collection
}

func NewAdminRep(ar *mongo.Collection) *AdminRep {
	return &AdminRep{rep: ar}
}
