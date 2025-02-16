package product_delivery

import "github.com/gin-gonic/gin"

type ProductDeliveryI interface {
	AddProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	EditProduct(c *gin.Context)
	FilterProducts(c *gin.Context)
	DeleteProduct(c *gin.Context)
	AddCategory(c *gin.Context)
}
