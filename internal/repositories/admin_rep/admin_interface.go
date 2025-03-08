package admin_rep

import "store/internal/entities"

type AdminRepI interface {
	AddAdmin(admin entities.Admin) error
}
