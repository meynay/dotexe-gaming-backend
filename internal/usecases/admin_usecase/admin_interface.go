package admin_usecase

import (
	"store/internal/entities"
	"time"
)

type AdminUsecaseI interface {
	//product
	AddProduct(product entities.Product) error
	EditProduct(product entities.Product) error
	DeleteProduct(id string) error

	//category
	AddCategory(category entities.Category) error
	EditCategory(category entities.Category) error
	DeleteCategory(id string) error

	//invoices
	GetInvoices(filter entities.InvoiceFilter) []entities.Invoice
	GetInvoice(id string) (entities.Invoice, error)
	ChangeInvoiceStatus(invoiceid string, status int) error

	//chart
	GetChartInfo(filter entities.ChartFilter) map[time.Time]int
}
