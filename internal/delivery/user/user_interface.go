package user_delivary

import "github.com/gin-gonic/gin"

type UserDeliveryI interface {
	FirstStep(c *gin.Context)
	LoginWithEmail(c *gin.Context)
	LoginWithPhone(c *gin.Context)
	SignupWithEmail(c *gin.Context)
	SignupWithPhone(c *gin.Context)
	RefreshToken(c *gin.Context)
}
