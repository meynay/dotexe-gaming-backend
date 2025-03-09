package invoice_rep

import (
	"context"
	"fmt"
	"log"
	"store/internal/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	log.Printf("invoice %s added with status %d", invoice.ID, invoice.OrderStatus)
	return nil
}

func (ir *InvoiceRep) GetInvoice(id primitive.ObjectID) (entities.Invoice, error) {
	invoice := entities.Invoice{}
	res := ir.rep.FindOne(context.TODO(), bson.M{"_id": id})
	if res.Err() != nil {
		return invoice, fmt.Errorf("couldn't find invoice with given id")
	}
	res.Decode(&invoice)
	return invoice, nil
}

func (ir *InvoiceRep) GetInvoices(userid primitive.ObjectID) ([]entities.Invoice, error) {
	invoices := []entities.Invoice{}
	res, err := ir.rep.Find(context.TODO(), bson.M{"user_id": userid})
	if err != nil {
		return invoices, fmt.Errorf("error getting invoices")
	}
	res.All(context.TODO(), &invoices)
	return invoices, nil
}

func (ir *InvoiceRep) GetAllInvoices() []entities.Invoice {
	invoices := []entities.Invoice{}
	res, err := ir.rep.Find(context.TODO(), bson.M{})
	if err != nil {
		return invoices
	}
	res.All(context.TODO(), &invoices)
	return invoices
}

func (ir *InvoiceRep) ChangeStatus(invoiceid primitive.ObjectID, status int) error {
	_, err := ir.rep.UpdateOne(context.TODO(), bson.M{"_id": invoiceid}, bson.M{"$set": bson.M{"order_status": status}})
	if err != nil {
		return fmt.Errorf("couldn't update status")
	}
	log.Printf("Updated invoice %s to status %d", invoiceid, status)
	return nil
}
