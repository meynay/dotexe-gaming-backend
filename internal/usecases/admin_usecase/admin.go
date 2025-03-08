package admin_usecase

import (
	"store/internal/repositories/admin_rep"
	"store/internal/repositories/category_rep"
	"store/internal/repositories/invoice_rep"
	"store/internal/repositories/product_rep"
)

type AdminUsecase struct {
	productrep  *product_rep.ProductRep
	categoryrep *category_rep.CategoryRep
	invoicerep  *invoice_rep.InvoiceRep
	adminrep    *admin_rep.AdminRep
}

func NewAdminUsecase(pr *product_rep.ProductRep, cr *category_rep.CategoryRep, ir *invoice_rep.InvoiceRep, ar *admin_rep.AdminRep) *AdminUsecase {
	return &AdminUsecase{productrep: pr, categoryrep: cr, invoicerep: ir, adminrep: ar}
}
