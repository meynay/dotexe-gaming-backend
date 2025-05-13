package cart_delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (cd *CartDelivery) GetInvoices(c *gin.Context) {
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	invoices := cd.usecase.GetInvoices(userID)
	c.JSON(http.StatusOK, invoices)
}

func (cd *CartDelivery) GetInvoice(c *gin.Context) {
	invoiceid, _ := strconv.Atoi(c.Param("invoiceid"))
	userid, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "userid not set"})
		return
	}
	userID, ok := userid.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error reading userid"})
	}
	invoice := cd.usecase.GetInvoice(userID, uint(invoiceid))
	c.JSON(http.StatusOK, invoice)
}
