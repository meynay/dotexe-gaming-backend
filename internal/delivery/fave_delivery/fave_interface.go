package fave_delivery

import "github.com/gin-gonic/gin"

type FaveDeliveryI interface {
	FaveProduct(c *gin.Context)
	UnfaveProduct(c *gin.Context)
	CheckFave(c *gin.Context)
	GetFaves(c *gin.Context)
}
