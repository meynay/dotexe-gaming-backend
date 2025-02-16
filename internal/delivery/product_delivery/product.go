package product_delivery

import (
	"store/internal/usecases/product_usecase"

	"github.com/gin-gonic/gin"
)

type ProductDelivery struct {
	pu *product_usecase.ProductUseCase
}

func NewProductDelivery(u *product_usecase.ProductUseCase) *ProductDelivery {
	return &ProductDelivery{pu: u}
}

func (pd *ProductDelivery) AddProduct(c *gin.Context) {

}

func (pd *ProductDelivery) GetProduct(c *gin.Context) {

}

func (pd *ProductDelivery) GetProducts(c *gin.Context) {

}

func (pd *ProductDelivery) EditProduct(c *gin.Context) {

}
