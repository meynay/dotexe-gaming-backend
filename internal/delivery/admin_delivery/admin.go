package admin_delivery

import "store/internal/usecases/admin_usecase"

type AdminDelivery struct {
	adminusecase *admin_usecase.AdminUsecase
}

func NewAdminDelivery(au *admin_usecase.AdminUsecase) *AdminDelivery {
	return &AdminDelivery{adminusecase: au}
}
