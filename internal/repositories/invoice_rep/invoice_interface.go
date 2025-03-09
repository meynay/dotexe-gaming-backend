package invoice_rep

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InvoiceRepI interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(id primitive.ObjectID) (entities.Invoice, error)
	GetInvoices(userid primitive.ObjectID) []entities.Invoice
	GetAllInvoices() []entities.Invoice
	ChangeStatus(invoiceid primitive.ObjectID, status int) error
}
