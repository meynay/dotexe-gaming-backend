package user_product_delivery

import "github.com/gin-gonic/gin"

type UserProductDeliveryI interface {
	FaveProduct(c *gin.Context)
	UnfaveProduct(c *gin.Context)
	CheckFave(c *gin.Context)
	GetFaves(c *gin.Context)
	CommentOnProduct(c *gin.Context)
	GetComments(c *gin.Context)
	RateProduct(c *gin.Context)
	GetRates(c *gin.Context)
	AddToCart(c *gin.Context)
	IsInCart(c *gin.Context)
}
