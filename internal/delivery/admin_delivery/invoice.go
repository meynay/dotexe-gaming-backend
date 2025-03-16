package admin_delivery

import (
	"net/http"
	"store/internal/entities"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ad *AdminDelivery) GetInvoices(c *gin.Context) {
	status, err := strconv.Atoi(c.Query("status"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status should be integer"})
		return
	}
	cts, err := strconv.Atoi(c.Query("counttoshow"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "counttoshow should be integer"})
		return
	}
	page, err := strconv.Atoi(c.Query("counttoshow"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page should be integer"})
		return
	}
	time1 := c.Query("from")
	time2 := c.Query("to")
	from, err := time.Parse(entities.TimeLayout, time1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad time format. should be like (2006-01-02)"})
		return
	}
	to, err := time.Parse(entities.TimeLayout, time2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad time format. should be like (2006-01-02)"})
		return
	}
	filter := entities.InvoiceFilter{
		Status:      status,
		From:        from,
		To:          to,
		CountToShow: cts,
		Page:        page,
	}
	invoices := ad.adminusecase.GetInvoices(filter)
	c.JSON(http.StatusOK, invoices)
}

func (ad *AdminDelivery) GetInvoice(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("invoiceid"))
	name, phone, invoice, err := ad.adminusecase.GetInvoice(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_name": name, "phone_number": phone, "order": invoice})
}

func (ad *AdminDelivery) ChangeInvoiceStatus(c *gin.Context) {
	id, _ := primitive.ObjectIDFromHex(c.Param("invoiceid"))
	status := c.Query("status")
	stat, _ := strconv.Atoi(status)
	err := ad.adminusecase.ChangeInvoiceStatus(id, stat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status changed for order"})
}
