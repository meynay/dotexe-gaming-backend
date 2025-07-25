package admin_delivery

import (
	"net/http"
	"store/internal/entities"
	"store/internal/usecases/admin_usecase"
	"store/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type AdminDelivery struct {
	adminusecase *admin_usecase.AdminUsecase
	generator    *jwt.JWTTokenHandler
}

func NewAdminDelivery(au *admin_usecase.AdminUsecase, gn *jwt.JWTTokenHandler) *AdminDelivery {
	return &AdminDelivery{adminusecase: au, generator: gn}
}

func (ad *AdminDelivery) Login(c *gin.Context) {
	inp := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	id, err := ad.adminusecase.Login(inp.Username, inp.Password)
	if err != nil {
		if id == 0 {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "wrong password"})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "no such username"})
		return
	}
	at, _, err := ad.generator.GenerateJWT(id, 60)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access token": at})
}

func (ad *AdminDelivery) AddAdmin(c *gin.Context) {
	inp := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.ShouldBindJSON(&inp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	err := ad.adminusecase.AddAdmin(inp.Username, inp.Password)
	if err != nil {
		if err.Error() == "this username exists" {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "username exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "admin added"})
}

func (ad *AdminDelivery) GetInfo(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
		return
	}
	admin, err := ad.adminusecase.GetInfo(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, admin)
}

func (ad *AdminDelivery) FillFields(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
		return
	}
	var admin entities.Admin
	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json format"})
		return
	}
	admin.ID = userID
	err := ad.adminusecase.FillFields(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "fields updated"})
}
