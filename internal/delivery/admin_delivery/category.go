package admin_delivery

import (
	"net/http"
	"store/internal/entities"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ad *AdminDelivery) AddCategory(c *gin.Context) {
	category := entities.Category{}
	if err := c.BindJSON(&category); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.AddCategory(category)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category added seccessfully"})
}

func (ad *AdminDelivery) EditCategory(c *gin.Context) {
	category := entities.Category{}
	if c.BindJSON(&category) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.EditCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category edited seccessfully"})
}

func (ad *AdminDelivery) DeleteCategory(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	err := ad.adminusecase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
}
