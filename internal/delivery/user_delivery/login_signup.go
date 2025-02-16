package user_delivery

import (
	"net/http"
	"store/internal/usecases/user_usecase"
	"store/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	uu        *user_usecase.UserUsecase
	generator jwt.JWTTokenHandler
}

func NewUserDelivary(uu *user_usecase.UserUsecase, j *jwt.JWTTokenHandler) *UserDelivery {
	return &UserDelivery{uu: uu}
}

func (d *UserDelivery) FirstStep(c *gin.Context) {
	input := struct {
		Inp string `json:"login"`
	}{}
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json", "status": 0})
		return
	}
	result, message := d.uu.FirstAttempt(input.Inp)
	out := struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}{}
	out.Status = result
	out.Message = message
	if result == 1 || result == 2 {
		c.JSON(http.StatusBadRequest, out)
		return
	}
	c.JSON(http.StatusOK, out)
}

func (d *UserDelivery) LoginWithEmail(c *gin.Context) {
	in := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := c.BindJSON(&in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}
	id, err := d.uu.LoginWithEmail(in.Email, in.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "wrong password"})
		return
	}
	at, rt, err := d.generator.GenerateJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	d.uu.SaveToken(id, rt)
	c.JSON(http.StatusOK, gin.H{
		"message":       "login successful",
		"access_token":  at,
		"refresh_token": rt,
	})
}

func (d *UserDelivery) LoginWithPhone(c *gin.Context) {
	in := struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}{}
	err := c.BindJSON(&in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}
	id, err := d.uu.LoginWithPhone(in.Phone, in.Code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "wrong code"})
		return
	}
	at, rt, err := d.generator.GenerateJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	d.uu.SaveToken(id, rt)
	c.JSON(http.StatusOK, gin.H{
		"message":       "login successful",
		"access_token":  at,
		"refresh_token": rt,
	})
}

func (d *UserDelivery) SignupWithEmail(c *gin.Context) {
	in := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	err := c.BindJSON(&in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}
	id, err := d.uu.SignupWithEmail(in.Email, in.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't register user"})
		return
	}
	at, rt, err := d.generator.GenerateJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	d.uu.SaveToken(id, rt)
	c.JSON(http.StatusOK, gin.H{
		"message":       "register successful",
		"access_token":  at,
		"refresh_token": rt,
	})
}

func (d *UserDelivery) SignupWithPhone(c *gin.Context) {
	in := struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}{}
	err := c.BindJSON(&in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid json"})
		return
	}
	id, err := d.uu.SignupWithPhone(in.Phone, in.Code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "wrong code"})
		return
	}
	at, rt, err := d.generator.GenerateJWT(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
	}
	d.uu.SaveToken(id, rt)
	c.JSON(http.StatusOK, gin.H{
		"message":       "register successful",
		"access_token":  at,
		"refresh_token": rt,
	})
}
