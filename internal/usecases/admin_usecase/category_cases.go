package admin_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (au *AdminUsecase) AddCategory(category entities.Category) error {
	return au.categoryrep.AddCategory(category)
}

func (au *AdminUsecase) EditCategory(category entities.Category) error {
	return au.categoryrep.EditCategory(category)
}

func (au *AdminUsecase) DeleteCategory(id primitive.ObjectID) error {
	return au.categoryrep.DeleteCategory(id)
}
