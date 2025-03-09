package admin_delivery

import (
	"net/http"
	"store/internal/entities"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ad *AdminDelivery) AddProduct(c *gin.Context) {
	product := entities.Product{}
	if c.BindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.AddProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added successfully"})
}

func (ad *AdminDelivery) EditProduct(c *gin.Context) {
	product := entities.Product{}
	if c.BindJSON(&product) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.EditProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "product edited successfully"})
}

func (ad *AdminDelivery) DeleteProduct(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err := ad.adminusecase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"message": "product deleted successfully"})
}
