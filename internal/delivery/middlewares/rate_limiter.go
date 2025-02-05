package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// sets a 2 second timeout for each request
func RateLimiter() gin.HandlerFunc {
	return timeout.New(timeout.WithTimeout(2*time.Second), timeout.WithHandler(func(c *gin.Context) {
		c.Next()
	}), timeout.WithResponse(func(c *gin.Context) {
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timeout, try again later"})
	}))
}
