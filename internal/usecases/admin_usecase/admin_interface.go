package admin_usecase

import (
	"store/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminUsecaseI interface {
	//admin
	AddAdmin(username, password string) error
	FillFields(admin entities.Admin) error
	Login(username, password string) (primitive.ObjectID, error)
	ForgetPassword1(phone string) error
	ForgetPassword2(phone, code string) error
	//product
	AddProduct(product entities.Product) error
	EditProduct(product entities.Product) error
	DeleteProduct(id primitive.ObjectID) error

	//category
	AddCategory(category entities.Category) error
	EditCategory(category entities.Category) error
	DeleteCategory(id primitive.ObjectID) error

	//invoices
	GetInvoices(filter entities.InvoiceFilter) []entities.Invoice
	GetInvoice(id primitive.ObjectID) (entities.Invoice, error)
	ChangeInvoiceStatus(invoiceid primitive.ObjectID, status int) error

	//chart
	GetChartInfo(filter entities.ChartFilter) map[time.Time]int
}
