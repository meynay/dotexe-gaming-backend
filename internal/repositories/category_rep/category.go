package category_rep

import (
	"errors"
	"fmt"
	"log"
	"store/internal/entities"

	"gorm.io/gorm"
)

type CategoryRep struct {
	rep *gorm.DB
}

func NewCategoryRep(crep *gorm.DB) *CategoryRep {
	return &CategoryRep{rep: crep}
}

func (cr *CategoryRep) AddCategory(c entities.Category) error {
	res := cr.rep.Create(&c)
	if res.Error != nil {
		return fmt.Errorf("couldn't insert category, %v", res.Error)
	}

	log.Printf("Category %s added with id %v\n", c.Name, c.ID)
	return nil
}

func (cr *CategoryRep) EditCategory(c entities.Category) error {
	result := cr.rep.Model(&entities.Category{}).
		Where("id = ?", c.ID).
		Updates(entities.Category{
			Name:     c.Name,
			Image:    c.Image,
			ParentID: c.ParentID,
		})

	if result.Error != nil {
		return fmt.Errorf("update failed: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("category not found")
	}
	return nil
}

func (cr *CategoryRep) DeleteCategory(id uint) error {
	var category entities.Category
	if err := cr.rep.First(&category, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("category with id %v not found", id)
		}
		return fmt.Errorf("database error: %v", err)
	}
	var childCount int64
	if err := cr.rep.Model(&entities.Category{}).Where("parent_id = ?", id).Count(&childCount).Error; err != nil {
		return fmt.Errorf("failed to check child categories: %v", err)
	}

	if childCount > 0 {
		return fmt.Errorf("cannot delete category with id %v - it has %d child categories", id, childCount)
	}

	err := cr.rep.Transaction(func(tx *gorm.DB) error {
		//soft delete
		if err := tx.Delete(&category).Error; err != nil {
			return err
		}
		//hard delete
		// if err := tx.Unscoped().Delete(&category).Error; err != nil {
		//     return err
		// }

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	log.Printf("Category with id %v deleted (Name: %s)", id, category.Name)
	return nil
}

func (cr *CategoryRep) GetCategory(id uint) (entities.Category, error) {
	var category entities.Category

	// Basic query without relationships
	// err := cr.db.First(&category, id).Error

	// Recommended: With relationships preloaded
	err := cr.rep.Preload("Parent").Preload("Children").First(&category, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Category{}, fmt.Errorf("category with id %v not found", id)
		}
		return entities.Category{}, fmt.Errorf("database error: %v", err)
	}

	return category, nil
}

func (cr *CategoryRep) GetCategories() ([]entities.Category, error) {
	var categories []entities.Category

	// Basic query without relationships
	// err := cr.db.Find(&categories).Error

	// Recommended: With hierarchical relationships preloaded
	err := cr.rep.Preload("Parent").
		Preload("Children").
		Order("name ASC"). // Example ordering
		Find(&categories).Error

	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %v", err)
	}

	return categories, nil
}

func (cr *CategoryRep) GetParents(ID uint) []uint {
	var parents []uint
	currentID := ID

	// Prevent infinite loops with a reasonable limit
	maxDepth := 10
	depth := 0

	for currentID != 0 && depth < maxDepth {
		var category entities.Category
		if err := cr.rep.Select("parent_id").
			First(&category, currentID).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil
			}
			return nil
		}

		if category.ParentID != nil {
			parents = append(parents, *category.ParentID)
			currentID = *category.ParentID
		} else {
			break
		}

		depth++
	}

	if depth == maxDepth {
		return nil
	}

	return parents
}
