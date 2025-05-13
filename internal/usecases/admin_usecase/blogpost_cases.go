package admin_usecase

import "store/internal/entities"

func (au *AdminUsecase) AddBlogPost(bp entities.BlogPost) error {
	return au.blogpostrep.AddBlogPost(bp)
}
func (au *AdminUsecase) EditBlogPost(bp entities.BlogPost) error {
	return au.blogpostrep.EditBlogPost(bp)
}
func (au *AdminUsecase) DeleteBlogPost(ID uint) error {
	return au.blogpostrep.DeleteBlogPost(ID)
}
func (au *AdminUsecase) AddComment(cm entities.BPComment) error {
	return au.blogpostrep.AddComment(cm)
}
