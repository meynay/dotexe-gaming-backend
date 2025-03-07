package user_product_delivery

import (
	"net/http"
	"store/internal/usecases/user_product_usecase"

	"github.com/gin-gonic/gin"
)

type UserProductDelivery struct {
	up *user_product_usecase.UserProductUsecase
}

func NewUserProductDelivery(upu *user_product_usecase.UserProductUsecase) *UserProductDelivery {
	return &UserProductDelivery{up: upu}
}

func (up *UserProductDelivery) FaveProduct(c *gin.Context) {
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
	err := up.up.FaveProduct(userID, productid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added to faves"})
}

func (up *UserProductDelivery) UnfaveProduct(c *gin.Context) {
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
	err := up.up.UnfaveProduct(userID, productid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"message": "product deleted from faves"})
}

func (up *UserProductDelivery) CheckFave(c *gin.Context) {
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
	err := up.up.CheckFave(userID, productid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product is not in user faves"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product is in user faves"})
}

func (up *UserProductDelivery) GetFaves(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	faves := up.up.GetFaves(userID)
	c.JSON(http.StatusOK, faves)
}

func (up *UserProductDelivery) CommentOnProduct(c *gin.Context) {

}

func (up *UserProductDelivery) GetComments(c *gin.Context) {
	productid := c.Param("id")
	out := up.up.GetComments(productid)
	c.JSON(http.StatusOK, out)
}

func (up *UserProductDelivery) RateProduct(c *gin.Context) {

}

func (up *UserProductDelivery) GetRates(c *gin.Context) {
	productid := c.Param("id")
	out := up.up.GetRates(productid)
	c.JSON(http.StatusOK, out)
}
