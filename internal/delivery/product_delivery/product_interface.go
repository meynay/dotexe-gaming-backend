package product_delivery

import "github.com/gin-gonic/gin"

type ProductDeliveryI interface {
	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	SearchQuery(c *gin.Context)
	GetCategories(c *gin.Context)
}
