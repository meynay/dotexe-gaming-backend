package comment_delivery

import "github.com/gin-gonic/gin"

type CommentDeliveryI interface {
	CommentOnProduct(c *gin.Context)
	GetComments(c *gin.Context)
}
