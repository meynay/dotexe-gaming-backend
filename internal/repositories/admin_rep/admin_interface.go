package admin_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminRepI interface {
	AddAdmin(username, password string) error
	GetName(id primitive.ObjectID) string
	FillFields(admin entities.Admin) error
	Login(username, password string) (primitive.ObjectID, error)
}
