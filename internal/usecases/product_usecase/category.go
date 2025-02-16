package product_usecase

import "store/internal/entities"

func (pu *ProductUseCase) AddCategory(c entities.Category) error {
	return pu.rep.AddCategory(c)
}
