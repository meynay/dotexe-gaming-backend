package user_delivery

import "github.com/gin-gonic/gin"

type UserDeliveryI interface {
	//user
	GetInfo(c *gin.Context)
	FillInfo(c *gin.Context)
	ResetPassword(c *gin.Context)

	//login
	FirstStep(c *gin.Context)
	LoginWithEmail(c *gin.Context)
	LoginWithPhone(c *gin.Context)
	SignupWithEmail(c *gin.Context)
	SignupWithPhone(c *gin.Context)
	RefreshToken(c *gin.Context)
}
