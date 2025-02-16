package blogpost_usecase

import "store/internal/entities"

type BlogPostUseCaseI interface {
	AddBlogPost(bp entities.BlogPost) error
	GetBlogPost(ID string) (entities.BlogPost, error)
	GetBlogPosts() ([]entities.BlogPost, error)
	EditBlogPost(bp entities.BlogPost) error
	DeleteBlogPost(ID string) error
	AddComment(cm entities.BPComment) error
	GetComments(ID string) ([]entities.BPComment, error)
}
