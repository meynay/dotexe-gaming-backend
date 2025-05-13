package admin_usecase

import (
	"fmt"
	"sort"
	"store/internal/entities"
)

func (au *AdminUsecase) GetChartInfo(filter entities.ChartFilter) map[string]int {
	data := make(map[string]int)
	invoices := au.invoicerep.GetAllInvoices()
	neededInvoices := []entities.Invoice{}
	for _, invoice := range invoices {
		if invoice.CreatedAt.Before(filter.To) && invoice.CreatedAt.After(filter.From) {
			neededInvoices = append(neededInvoices, invoice)
		}
	}
	for !filter.From.Equal(filter.To) {
		y, m, d := filter.From.Date()
		str := fmt.Sprintf("%d-%v-%d", y, m, d)
		data[str] = 0
		filter.From = filter.From.AddDate(0, 0, 1)
	}
	sort.Slice(neededInvoices, func(i, j int) bool {
		return neededInvoices[i].CreatedAt.Before(neededInvoices[j].CreatedAt)
	})
	for _, invoice := range neededInvoices {
		y, m, d := invoice.CreatedAt.Date()
		str := fmt.Sprintf("%d-%v-%d", y, m, d)
		if filter.ShowType == entities.OrdersCount {
			data[str] += 1
		} else if filter.ShowType == entities.TotalPrice {
			data[str] += invoice.TotalPrice
		} else {
			itemsCount := 0
			for _, item := range invoice.Items {
				itemsCount += item.Count
			}
			data[str] += itemsCount
		}
	}
	return data
}
