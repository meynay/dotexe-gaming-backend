package middlewares

import (
	"net/http"
	"store/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	j *jwt.JWTTokenHandler
}

func NewAuth(j *jwt.JWTTokenHandler) *Auth {
	return &Auth{j: j}
}

func (a *Auth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//gets token from request header and checks if it's empty
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}
		//validates token and gets phone_number out of it
		id, err := a.j.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		//set phone number and sends to endpoints
		c.Set("id", id)
		c.Next()
	}
}
