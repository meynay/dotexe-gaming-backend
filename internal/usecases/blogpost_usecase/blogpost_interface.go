package blogpost_usecase

import "store/internal/entities"

type BlogPostUseCaseI interface {
	GetBlogPost(ID string) (entities.BlogPost, error)
	GetBlogPosts() ([]entities.BlogPost, error)

	AddComment(cm entities.BPComment) error
	GetComments(ID string) ([]entities.BPComment, error)
}
