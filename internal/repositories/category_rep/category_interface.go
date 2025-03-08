package category_rep

import "store/internal/entities"

type CategoryRepI interface {
	AddCategory(c entities.Category) error
	EditCategory(c entities.Category) error
	DeleteCategory(id string) error
	GetCategory(id string) (entities.Category, error)
	GetCategories() []entities.Category
}
