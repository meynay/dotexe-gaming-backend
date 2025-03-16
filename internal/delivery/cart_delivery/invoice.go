package cart_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cd *CartDelivery) GetInvoices(c *gin.Context) {
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
	invoices := cd.usecase.GetInvoices(userd)
	c.JSON(http.StatusOK, invoices)
}

func (cd *CartDelivery) GetInvoice(c *gin.Context) {
	invoiceid, _ := primitive.ObjectIDFromHex(c.Param("invoiceid"))
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
	cd.usecase.GetInvoice(userd, invoiceid)
}
