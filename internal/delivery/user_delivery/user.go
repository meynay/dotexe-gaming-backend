package user_delivery

import (
	"net/http"
	"store/internal/entities"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ud *UserDelivery) GetInfo(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	userd, _ := primitive.ObjectIDFromHex(userID)
	user, err := ud.uu.GetInfo(userd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (ud *UserDelivery) FillInfo(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	userd, _ := primitive.ObjectIDFromHex(userID)
	var user entities.User
	if c.BindJSON(&user) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	user.ID = userd
	err := ud.uu.FillInfo(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user info updated"})
}

func (ud *UserDelivery) ResetPassword(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	userd, _ := primitive.ObjectIDFromHex(userID)
	pass := struct {
		Password string `json:"password"`
	}{}
	if c.Bind(&pass) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
		return
	}
	err := ud.uu.ResetPassword(userd, pass.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password updated"})
}
