package cart_usecase

import (
	"store/internal/entities"
)

func (cu *CartUsecase) GetInvoices(userid uint) []entities.Invoice {
	invoices, err := cu.invoicerep.GetInvoices(userid)
	if err != nil {
		return []entities.Invoice{}
	}
	return invoices
}

func (cu *CartUsecase) GetInvoice(userid, invoiceid uint) entities.Invoice {
	invoice, err := cu.invoicerep.GetInvoice(invoiceid)
	if err != nil || invoice.UserID != userid {
		return entities.Invoice{}
	}
	return invoice
}
