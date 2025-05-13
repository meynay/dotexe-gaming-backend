package blogpost_rep

import (
	"fmt"
	"store/internal/entities"

	"gorm.io/gorm"
)

type BlogPostRep struct {
	bpdb *gorm.DB
}

func NewBlogPostRep(bp *gorm.DB) *BlogPostRep {
	return &BlogPostRep{bpdb: bp}
}

func (bp *BlogPostRep) AddBlogPost(b entities.BlogPost) error {
	tx := bp.bpdb.Create(&b)
	if tx.Error != nil {
		return fmt.Errorf("couldn't insert blogpost")
	}
	return nil
}

func (bp *BlogPostRep) GetBlogPost(ID uint) (entities.BlogPost, error) {
	var blogpost entities.BlogPost
	res := bp.bpdb.First(&blogpost, ID)
	if res.Error != nil {
		return entities.BlogPost{}, fmt.Errorf("no documents found")
	}
	return blogpost, nil
}

func (bp *BlogPostRep) GetBlogPosts() ([]entities.BlogPost, error) {
	var posts []entities.BlogPost
	tx := bp.bpdb.Find(&posts)
	if tx.Error != nil {
		return posts, tx.Error
	}
	return posts, nil
}

func (bp *BlogPostRep) EditBlogPost(b entities.BlogPost) error {
	tx := bp.bpdb.Save(&b)
	if tx.Error != nil {
		return fmt.Errorf("couldn't update blogpost")
	}
	return nil
}

func (bp *BlogPostRep) DeleteBlogPost(ID uint) error {
	tx := bp.bpdb.Delete(&entities.BlogPost{}, ID)
	if tx.Error != nil {
		return fmt.Errorf("couldn't delete blogpost")
	}
	return nil
}
