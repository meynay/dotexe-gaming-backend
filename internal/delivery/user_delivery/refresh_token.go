package user_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	id, _ := primitive.ObjectIDFromHex(ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}
	if !d.uu.TokenExists(id, request.RefreshToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	//generates new access token
	newAccessToken, _, err := d.generator.GenerateJWT(id, 15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
