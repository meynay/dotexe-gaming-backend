package admin_usecase

import "store/internal/entities"

func (au *AdminUsecase) AddCategory(category entities.Category) error {
	return au.categoryrep.AddCategory(category)
}

func (au *AdminUsecase) EditCategory(category entities.Category) error {
	return au.categoryrep.EditCategory(category)
}

func (au *AdminUsecase) DeleteCategory(id string) error {
	return au.categoryrep.DeleteCategory(id)
}
