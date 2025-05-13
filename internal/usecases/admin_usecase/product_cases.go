package admin_usecase

import (
	"store/internal/entities"
)

func (au *AdminUsecase) AddProduct(product entities.Product) error {
	return au.productrep.AddProduct(product)
}

func (au *AdminUsecase) EditProduct(product entities.Product) error {
	return au.productrep.EditProduct(product)
}

func (au *AdminUsecase) DeleteProduct(id uint) error {
	return au.productrep.DeleteProduct(id)
}

func (au *AdminUsecase) GetActiveProductsCount() int {
	count := 0
	products, _ := au.productrep.GetProducts()
	for _, pr := range products {
		if pr.Stock > 0 {
			count++
		}
	}
	return count
}
