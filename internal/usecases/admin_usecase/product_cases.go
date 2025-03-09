package admin_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (au *AdminUsecase) AddProduct(product entities.Product) error {
	return au.productrep.AddProduct(product)
}

func (au *AdminUsecase) EditProduct(product entities.Product) error {
	return au.productrep.EditProduct(product)
}

func (au *AdminUsecase) DeleteProduct(id primitive.ObjectID) error {
	return au.productrep.DeleteProduct(id)
}
