package admin_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/admin_rep"
	"store/internal/repositories/blogpost_rep"
	"store/internal/repositories/category_rep"
	"store/internal/repositories/invoice_rep"
	"store/internal/repositories/product_rep"
	"store/internal/repositories/user_rep"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminUsecase struct {
	productrep  *product_rep.ProductRep
	categoryrep *category_rep.CategoryRep
	invoicerep  *invoice_rep.InvoiceRep
	adminrep    *admin_rep.AdminRep
	userrep     *user_rep.UserRepository
	blogpostrep *blogpost_rep.BlogPostRep
}

func NewAdminUsecase(pr *product_rep.ProductRep, cr *category_rep.CategoryRep, ir *invoice_rep.InvoiceRep, ar *admin_rep.AdminRep, ur *user_rep.UserRepository, bpr *blogpost_rep.BlogPostRep) *AdminUsecase {
	return &AdminUsecase{productrep: pr, categoryrep: cr, invoicerep: ir, adminrep: ar, userrep: ur, blogpostrep: bpr}
}

func (au *AdminUsecase) AddAdmin(username, password string) error {
	return au.adminrep.AddAdmin(username, password)
}

func (au *AdminUsecase) GetInfo(adminID primitive.ObjectID) (entities.Admin, error) {
	return au.adminrep.GetInfo(adminID)
}

func (au *AdminUsecase) FillFields(admin entities.Admin) error {
	return au.adminrep.FillFields(admin)
}

func (au *AdminUsecase) Login(username, password string) (primitive.ObjectID, error) {
	return au.adminrep.Login(username, password)
}
