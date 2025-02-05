package middlewares

import (
	"net/http"
	"store/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//gets token from request header and checks if it's empty
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		//validates token and gets phone_number out of it
		phoneNumber, err := jwt.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		//set phone number and sends to endpoints
		c.Set("phone_number", phoneNumber)
		c.Next()
	}
}
