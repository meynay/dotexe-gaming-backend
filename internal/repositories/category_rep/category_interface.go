package category_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryRepI interface {
	AddCategory(c entities.Category) error
	EditCategory(c entities.Category) error
	DeleteCategory(id primitive.ObjectID) error
	GetCategory(id primitive.ObjectID) (entities.Category, error)
	GetCategories() []entities.Category
	GetParents(ID primitive.ObjectID) []primitive.ObjectID
}
