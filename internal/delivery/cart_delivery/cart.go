package cart_delivery

import (
	"net/http"
	"store/internal/usecases/cart_usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartDelivery struct {
	usecase *cart_usecase.CartUsecase
}

func NewCartDelivery(cu *cart_usecase.CartUsecase) *CartDelivery {
	return &CartDelivery{usecase: cu}
}

func (cd *CartDelivery) AddToCart(c *gin.Context) {
	productid, _ := primitive.ObjectIDFromHex(c.Param("id"))
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
	err := cd.usecase.AddToCart(productid, userd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "product added to cart"})
}

func (cd *CartDelivery) ChangeCountInCart(c *gin.Context) {
	productid, _ := primitive.ObjectIDFromHex(c.Param("id"))
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
	change := c.Param("change")
	if change == "inc" {
		err := cd.usecase.IncreaseInCart(productid, userd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "increased product count in cart"})
		return
	}
	count, err := cd.usecase.IsInCart(productid, userd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
		return
	}
	if count == 1 {
		err := cd.usecase.DeleteFromCart(productid, userd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted product from cart"})
		return
	}
	err = cd.usecase.DecreaseInCart(productid, userd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "decreased product count in cart"})
}

func (cd *CartDelivery) IsInCart(c *gin.Context) {
	productid, _ := primitive.ObjectIDFromHex(c.Param("id"))
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
	count, err := cd.usecase.IsInCart(productid, userd)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product is not in cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"count": count})
}

func (cd *CartDelivery) GetCart(c *gin.Context) {
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
	cart := cd.usecase.GetCart(userd)
	c.JSON(http.StatusOK, cart)
}

func (cd *CartDelivery) FinalizeCart(c *gin.Context) {

}
