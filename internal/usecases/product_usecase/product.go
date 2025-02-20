package product_usecase

import (
	"sort"
	"store/internal/entities"
	"store/internal/repositories/product_rep"
	"store/pkg"
)

type ProductUseCase struct {
	rep *product_rep.ProductRep
}

func NewProductUseCase(r *product_rep.ProductRep) *ProductUseCase {
	return &ProductUseCase{rep: r}
}

func (pu *ProductUseCase) AddProduct(p entities.Product) error {
	return pu.rep.AddProduct(p)
}

func (pu *ProductUseCase) GetProduct(ID string) (entities.Product, error) {
	err := pu.rep.AddViewToProduct(ID)
	if err != nil {
		return entities.Product{}, err
	}
	return pu.rep.GetProduct(ID)
}

func (pu *ProductUseCase) GetProducts() ([]entities.ProductLess, error) {
	products, err := pu.rep.GetProducts()
	var p []entities.ProductLess
	if err != nil {
		return p, err
	}
	for _, product := range products {
		var pr = entities.ProductLess{
			ID:    product.ID,
			Image: product.Image,
			Name:  product.Name,
			Price: product.Price,
			Off:   product.Off,
		}
		pr.Category, err = pu.rep.GetCategoryName(product.CategoryID)
		if err != nil {
			return p, err
		}
		p = append(p, pr)
	}
	return p, nil
}

func (pu *ProductUseCase) EditProduct(p entities.Product) error {
	return pu.rep.EditProduct(p)
}

func (pu *ProductUseCase) DeleteProduct(ID string) error {
	return pu.rep.DeleteProduct(ID)
}

func (pu *ProductUseCase) FilterProducts(filter entities.Filter) ([]entities.ProductLess, int, error) {
	products, err := pu.rep.GetProducts()
	var p []entities.ProductLess
	if err != nil {
		return p, 0, err
	}
	if filter.CategoryID != "" {
		newproducts := []entities.Product{}
		for _, product := range products {
			if product.CategoryID == filter.CategoryID || pkg.Exists(filter.CategoryID, pu.rep.GetParents(product.CategoryID)) {
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
			ID:    ps[start].Pr.ID,
			Name:  ps[start].Pr.Name,
			Image: ps[start].Pr.Image,
			Price: ps[start].Pr.Price,
			Off:   ps[start].Pr.Off,
		}
		out.Category, _ = pu.rep.GetCategoryName(ps[start].Pr.CategoryID)
		p = append(p, out)
		start++
	}
	return p, pages, nil
}
