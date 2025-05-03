package admin_usecase

import (
	"sort"
	"store/internal/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (au *AdminUsecase) GetInvoices(filter entities.InvoiceFilter) []entities.Invoice {
	temp := au.invoicerep.GetAllInvoices()
	invoices := []entities.Invoice{}
	for _, invoice := range temp {
		if filter.Status != entities.All && invoice.OrderStatus != filter.Status {
			continue
		}
		if filter.From.After(invoice.InvoiceDate) || filter.To.Before(invoice.InvoiceDate) {
			continue
		}
		invoices = append(invoices, invoice)
	}
	sort.Slice(invoices, func(i, j int) bool {
		return invoices[i].InvoiceDate.After(invoices[j].InvoiceDate)
	})
	from := (filter.Page - 1) * filter.CountToShow
	to := from + filter.CountToShow
	if len(invoices) == 0 {
		return invoices
	} else if to > len(invoices)-1 {
		return invoices[from:]
	}
	return invoices[from:to]
}

func (au *AdminUsecase) GetInvoice(id primitive.ObjectID) (string, string, entities.Invoice, error) {
	invoice, err := au.invoicerep.GetInvoice(id)
	if err != nil {
		return "", "", invoice, err
	}
	return au.userrep.GetUsername(invoice.UserID), au.userrep.GetPhoneNumber(invoice.UserID), invoice, nil
}

func (au *AdminUsecase) ChangeInvoiceStatus(invoiceid primitive.ObjectID, status int) error {
	return au.invoicerep.ChangeStatus(invoiceid, status)
}

func (au *AdminUsecase) GetNewInvoicesCount() int {
	return len(au.GetInvoices(entities.InvoiceFilter{
		From:        time.Now().AddDate(0, 0, -8),
		To:          time.Now().AddDate(0, 0, 1),
		CountToShow: 99999,
		Page:        1,
		Status:      entities.All,
	}))
}

func (au *AdminUsecase) GetMonthlySalesPrice() int {
	invoices := au.GetInvoices(entities.InvoiceFilter{
		From:        time.Now().AddDate(0, -1, -1),
		To:          time.Now().AddDate(0, 0, 1),
		CountToShow: 99999,
		Page:        1,
		Status:      entities.All,
	})
	price := 0
	for _, invoice := range invoices {
		price += invoice.TotalPrice
	}
	return price
}
