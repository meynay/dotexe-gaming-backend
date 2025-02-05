package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//gets api key
		apiKey := c.GetHeader("x-api-key")

		//checks if api key is empty
		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing API Key"})
			c.Abort()
			return
		}

		//checks if api key matches
		if apiKey != os.Getenv("API_KEY") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized API Key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
