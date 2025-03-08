package invoice_rep

import "store/internal/entities"

type InvoiceRepI interface {
	AddInvoice(invoice entities.Invoice) error
	GetInvoice(id string) (entities.Invoice, error)
	GetInvoices(userid string) []entities.Invoice
	GetAllInvoices() []entities.Invoice
	ChangeStatus(invoiceid string, status int) error
}
