package blogpost_usecase

import (
	"store/internal/entities"
	"store/internal/repositories/blogpost_rep"
	"store/pkg"
)

type BlogPostUseCase struct {
	bpr *blogpost_rep.BlogPostRep
}

func NewBlogPostUseCase(bp *blogpost_rep.BlogPostRep) *BlogPostUseCase {
	return &BlogPostUseCase{bpr: bp}
}

func (b *BlogPostUseCase) GetBlogPost(ID string) (entities.BlogPostR, error) {
	_, err := b.bpr.GetBlogPost(ID)
	if err != nil {
		return entities.BlogPostR{}, err
	}
	return entities.BlogPostR{}, nil
}

func (b *BlogPostUseCase) GetBlogPosts(filter entities.BPFilter) ([]entities.MiniBP, error) {
	blogposts, _ := b.bpr.GetBlogPosts()
	chosenones := []entities.MiniBP{}
	for _, blogpost := range blogposts {
		if pkg.Exists(blogpost.Category_id, filter.Categories) {
			if pkg.CalculateScore(filter.Query, blogpost.Title) > 0.7 {
				chosenones = append(chosenones, entities.MiniBP{
					ID:        blogpost.ID,
					Title:     blogpost.Title,
					Image:     blogpost.Image,
					Likes:     blogpost.Likes,
					Dislikes:  blogpost.Dislikes,
					UpdatedAt: blogpost.UpdatedAt,
					Author:    "sd",
					Category:  "md",
				})
			}
		}
	}
	return chosenones, nil
}
