package user_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *UserDelivery) RefreshToken(c *gin.Context) {
	//gets the refresh token
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//validates refresh token and takes phone number out of it
	ID, err := d.generator.ValidateJWT(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	// ID, err := strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
	// 	return
	// }
	if !d.uu.TokenExists(uint(ID), request.RefreshToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	//generates new access token
	newAccessToken, _, err := d.generator.GenerateJWT(uint(ID), 15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
