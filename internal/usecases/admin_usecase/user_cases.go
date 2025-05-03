package admin_usecase

import (
	"store/internal/entities"
	"store/pkg"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (au *AdminUsecase) GetActiveUsersCount() int {
	return len(au.GetActiveUsers())
}

func (au *AdminUsecase) GetActiveUsers() []entities.User {
	ids := []primitive.ObjectID{}
	invoices := au.GetInvoices(entities.InvoiceFilter{
		Status:      entities.All,
		CountToShow: 999999,
		Page:        1,
		From:        time.Now().AddDate(0, -1, -1),
		To:          time.Now().AddDate(0, 0, 1),
	})
	for _, invoice := range invoices {
		if !pkg.Exists(invoice.UserID, ids) {
			ids = append(ids, invoice.UserID)
		}
	}
	users := []entities.User{}
	for _, id := range ids {
		user, _ := au.GetUser(id)
		users = append(users, user)
	}
	return users
}

func (au *AdminUsecase) GetUser(ID primitive.ObjectID) (entities.User, error) {
	user, err := au.userrep.GetInfo(ID)
	if err != nil {
		return user, err
	}
	user.Password = "encrypted"
	user.Faves = []primitive.ObjectID{}
	user.Cart = []entities.Item{}
	return user, nil
}
