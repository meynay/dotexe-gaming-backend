package admin_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ad *AdminDelivery) GetUser(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("userid"))
	user, err := ad.adminusecase.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ad *AdminDelivery) GetActiveUsers(c *gin.Context) {
	c.JSON(http.StatusOK, ad.adminusecase.GetActiveUsers())
}

func (ad *AdminDelivery) GetActiveUsersCount(c *gin.Context) {
	c.JSON(http.StatusOK, ad.adminusecase.GetActiveUsersCount())
}
