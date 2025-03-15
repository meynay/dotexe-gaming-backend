package admin_delivery

import "github.com/gin-gonic/gin"

type AdminDeliveryI interface {
	Login(c *gin.Context)
	AddAdmin(c *gin.Context)
	FillFields(c *gin.Context)
	//product
	AddProduct(c *gin.Context)
	EditProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)

	//category
	AddCategory(c *gin.Context)
	EditCategory(c *gin.Context)
	DeleteCategory(c *gin.Context)

	//invoices
	GetInvoices(c *gin.Context)
	GetInvocie(c *gin.Context)
	ChangeInvoiceStatus(c *gin.Context)

	//chart
	GetChart(c *gin.Context)
}
