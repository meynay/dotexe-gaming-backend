package fave_delivery

import (
	"net/http"
	"store/internal/usecases/fave_usecase"

	"github.com/gin-gonic/gin"
)

type FaveDelivery struct {
	faveusecase *fave_usecase.FaveUsecase
}

func NewFaveDelivery(fu *fave_usecase.FaveUsecase) *FaveDelivery {
	return &FaveDelivery{faveusecase: fu}
}

func (fd *FaveDelivery) FaveProduct(c *gin.Context) {
	productid := c.Param("id")
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	err := fd.faveusecase.FaveProduct(userID, productid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added to faves"})
}

func (fd *FaveDelivery) UnfaveProduct(c *gin.Context) {
	productid := c.Param("id")
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	err := fd.faveusecase.UnfaveProduct(userID, productid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted from faves"})
}

func (fd *FaveDelivery) CheckFave(c *gin.Context) {
	productid := c.Param("id")
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	err := fd.faveusecase.CheckFave(userID, productid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product is not in user faves"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product is in user faves"})
}

func (fd *FaveDelivery) GetFaves(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	faves := fd.faveusecase.GetFaves(userID)
	c.JSON(http.StatusOK, faves)
}
