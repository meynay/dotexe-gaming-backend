package admin_delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ad *AdminDelivery) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("userid"))
	user, err := ad.adminusecase.GetUser(uint(id))
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
