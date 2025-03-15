package cart_delivery

import "github.com/gin-gonic/gin"

type CartDeliveryI interface {
	AddToCart(c *gin.Context)
	ChangeCountInCart(c *gin.Context)
	IsInCart(c *gin.Context)
	GetCart(c *gin.Context)
	FinalizeCart(c *gin.Context)
	GetInvoices(c *gin.Context)
	GetInvoice(c *gin.Context)
}
