package blogpost_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/blogpost_rep"
)

type BlogPostUseCase struct {
	bpr *blogpost_rep.BlogPostRep
}

func NewBlogPostUseCase(bp *blogpost_rep.BlogPostRep) *BlogPostUseCase {
	return &BlogPostUseCase{bpr: bp}
}

func (b *BlogPostUseCase) AddBlogPost(bp entities.BlogPost) error {
	return b.bpr.AddBlogPost(bp)
}
func (b *BlogPostUseCase) GetBlogPost(ID string) (entities.BlogPostR, error) {
	_, err := b.bpr.GetBlogPost(ID)
	if err != nil {
		return entities.BlogPostR{}, err
	}
	return entities.BlogPostR{}, nil
}

func (b *BlogPostUseCase) GetBlogPosts() ([]entities.BlogPost, error)
func (b *BlogPostUseCase) EditBlogPost(bp entities.BlogPost) error
func (b *BlogPostUseCase) DeleteBlogPost(ID string) error
func (b *BlogPostUseCase) AddComment(cm entities.BPComment) error
func (b *BlogPostUseCase) GetComments(ID string) ([]entities.BPComment, error)
