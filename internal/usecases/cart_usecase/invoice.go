package cart_usecase

import (
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (cu *CartUsecase) GetInvoices(userid primitive.ObjectID) []entities.Invoice {
	invoices, err := cu.invoicerep.GetInvoices(userid)
	if err != nil {
		return []entities.Invoice{}
	}
	return invoices
}

func (cu *CartUsecase) GetInvoice(userid, invoiceid primitive.ObjectID) entities.Invoice {
	invoice, err := cu.invoicerep.GetInvoice(invoiceid)
	if err != nil || invoice.UserID != userid {
		return entities.Invoice{}
	}
	return invoice
}
