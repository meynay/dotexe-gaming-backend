package invoice_rep

import (
	"fmt"
	"log"
	"store/internal/entities"

	"gorm.io/gorm"
)

type InvoiceRep struct {
	rep *gorm.DB
}

func NewInvoiceRep(irep *gorm.DB) *InvoiceRep {
	return &InvoiceRep{rep: irep}
}

func (ir *InvoiceRep) AddInvoice(invoice entities.Invoice) error {
	tx := ir.rep.Create(&invoice)
	if tx.Error != nil {
		return fmt.Errorf("couldn't insert invoice")
	}
	log.Printf("سفارش %d با وضعیت %d اضافه شد", invoice.ID, invoice.OrderStatus)
	return nil
}

func (ir *InvoiceRep) GetInvoice(id uint) (entities.Invoice, error) {
	invoice := entities.Invoice{}
	res := ir.rep.First(&invoice, id)
	if res.Error != nil {
		return invoice, fmt.Errorf("couldn't find invoice with given id")
	}
	return invoice, nil
}

func (ir *InvoiceRep) GetInvoices(userid uint) ([]entities.Invoice, error) {
	invoices := []entities.Invoice{}
	tx := ir.rep.Where("user_id = ?", userid).Find(&invoices)
	if tx.Error != nil {
		return invoices, fmt.Errorf("error getting invoices")
	}
	return invoices, nil
}

func (ir *InvoiceRep) GetAllInvoices() []entities.Invoice {
	invoices := []entities.Invoice{}
	tx := ir.rep.Find(&invoices)
	if tx.Error != nil {
		return []entities.Invoice{}
	}
	return invoices
}

func (ir *InvoiceRep) ChangeStatus(invoiceid uint, status int) error {
	tx := ir.rep.Model(entities.Invoice{}).Where("id = ?", invoiceid).Update("status = ?", status)
	if tx.Error != nil {
		return fmt.Errorf("couldn't update status")
	}
	log.Printf("سفارش به شماره %d به وضعیت %d تغییر یافت.", invoiceid, status)
	return nil
}
