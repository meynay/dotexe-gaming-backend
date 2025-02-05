package loginsignup

import (
	"net/http"
	"store/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	//gets the refresh token
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//validates refresh token and takes phone number out of it
	phoneNumber, err := jwt.ValidateJWT(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	//generates new access token
	newAccessToken, _, err := jwt.GenerateJWT(phoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
	})
}
