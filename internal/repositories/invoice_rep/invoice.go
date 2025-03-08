package invoice_rep

import (
	"context"
	"fmt"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type InvoiceRep struct {
	rep *mongo.Collection
}

func NewInvoiceRep(irep *mongo.Collection) *InvoiceRep {
	return &InvoiceRep{rep: irep}
}

func (ir *InvoiceRep) AddInvoice(invoice entities.Invoice) error {
	_, err := ir.rep.InsertOne(context.TODO(), invoice)
	if err != nil {
		return fmt.Errorf("couldn't insert invoice")
	}
	return nil
}

func (ir *InvoiceRep) GetInvoices(userid string) ([]entities.Invoice, error) {
	invoices := []entities.Invoice{}
	res, err := ir.rep.Find(context.TODO(), bson.M{"user_id": userid})
	if err != nil {
		return invoices, fmt.Errorf("error getting invoices")
	}
	res.Decode(&invoices)
	return invoices, nil
}
