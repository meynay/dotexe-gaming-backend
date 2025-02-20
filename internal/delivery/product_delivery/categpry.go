package product_delivery

import (
	"net/http"
	"store/internal/entities"

	"github.com/gin-gonic/gin"
)

func (pd *ProductDelivery) AddCategory(c *gin.Context) {
	category := entities.Category{}
	if c.BindJSON(&category) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad json format"})
		return
	}
	if pd.pu.AddCategory(category) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't add categpry"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category added successfully"})
}
