package product_usecase

import (
	"sort"
	"store/internal/entities"
	"store/internal/repositories/category_rep"
	"store/internal/repositories/product_rep"
	"store/pkg"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase struct {
	rep         *product_rep.ProductRep
	categoryrep *category_rep.CategoryRep
}

func NewProductUseCase(r *product_rep.ProductRep, cr *category_rep.CategoryRep) *ProductUseCase {
	return &ProductUseCase{rep: r, categoryrep: cr}
}

func (pu *ProductUseCase) GetProduct(ID primitive.ObjectID) (entities.Product, error) {
	err := pu.rep.AddViewToProduct(ID)
	if err != nil {
		return entities.Product{}, err
	}
	return pu.rep.GetProduct(ID)
}

func (pu *ProductUseCase) GetProducts(query string) ([]entities.ProductLess, []entities.Category, error) {
	var prdcts []entities.ProductLess
	var ctgrs []entities.Category
	products, err := pu.rep.GetProducts()
	categories := pu.categoryrep.GetCategories()
	if err != nil {
		return prdcts, ctgrs, err
	}
	for _, product := range products {
		if strings.Contains(product.Name, query) {
			prdcts = append(prdcts, entities.ProductLess{
				ID:          product.ID,
				Name:        product.Name,
				Image:       product.Image,
				Description: product.Description,
				Rating:      product.Rating,
				RateCount:   product.RateCount,
			})
		}
	}
	for _, category := range categories {
		if strings.Contains(category.Name, query) {
			ctgrs = append(ctgrs, category)
		}
	}
	return prdcts, ctgrs, nil
}

func (pu *ProductUseCase) FilterProducts(filter entities.Filter) ([]entities.ProductLess, int, error) {
	products, err := pu.rep.GetProducts()
	var p []entities.ProductLess
	if err != nil {
		return p, 0, err
	}
	z, _ := primitive.ObjectIDFromHex("0")
	if filter.CategoryID != z {
		newproducts := []entities.Product{}
		for _, product := range products {
			if product.CategoryID == filter.CategoryID || pkg.Exists(filter.CategoryID, pu.categoryrep.GetParents(product.CategoryID)) {
				newproducts = append(newproducts, product)
			}
		}
		products = newproducts
	}
	ps := []entities.PScore{}
	if filter.Query == "" {
		for _, p := range products {
			ps = append(ps, entities.PScore{Pr: p, Score: 1})
		}
	} else {
		for _, p := range products {
			score := pkg.CalculateScore(p.Name, filter.Query)
			if score > 0.7 {
				ps = append(ps, entities.PScore{Pr: p, Score: score})
			}

		}
	}
	switch filter.Order {
	case entities.CheapToExpensive:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Pr.Price < ps[j].Pr.Price
		})
	case entities.ExpensiveToCheap:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Pr.Price > ps[j].Pr.Price
		})
	case entities.MostOffToLeast:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Pr.Off > ps[j].Pr.Off
		})
	case entities.Newest:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Pr.AddedAt.After(ps[j].Pr.AddedAt)
		})
	case entities.MostPurchased:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Pr.PurchaseCount > ps[j].Pr.PurchaseCount
		})
	case entities.MostRelevant:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Score > ps[j].Score
		})
	case entities.MostViewed:
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].Pr.Views > ps[j].Pr.Views
		})
	case entities.MostRate:
	}
	pages := len(ps) / filter.NumberOfItems
	start := filter.NumberOfItems * (filter.Page - 1)
	end := start + filter.NumberOfItems
	for start < end && (start >= 0 && start < len(ps)) {
		out := entities.ProductLess{
			ID:            ps[start].Pr.ID,
			Name:          ps[start].Pr.Name,
			Image:         ps[start].Pr.Image,
			Price:         ps[start].Pr.Price,
			Off:           ps[start].Pr.Off,
			Description:   ps[start].Pr.Description,
			Rating:        ps[start].Pr.Rating,
			RateCount:     ps[start].Pr.RateCount,
			PurchaseCount: ps[start].Pr.PurchaseCount,
			Views:         ps[start].Pr.Views,
		}
		cat, _ := pu.categoryrep.GetCategory(ps[start].Pr.CategoryID)
		out.Category = cat.Name
		p = append(p, out)
		start++
	}
	return p, pages, nil
}

func (pu *ProductUseCase) GetCategories() []entities.Category {
	return pu.categoryrep.GetCategories()
}
