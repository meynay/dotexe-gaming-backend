package product_rep

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"store/internal/entities"

	"gorm.io/gorm"
)

type ProductRep struct {
	prdb *gorm.DB
}

func NewProductRep(pr *gorm.DB) *ProductRep {
	return &ProductRep{prdb: pr}
}

func (pr *ProductRep) AddProduct(p entities.Product) error {
	tagsJSON, err := json.Marshal(p.Tags)
	if err != nil {
		return err
	}
	p.Tags = nil
	res := pr.prdb.Create(&p)

	if res.Error != nil {
		return fmt.Errorf("couldn't add product, %v", res.Error)
	}
	pr.prdb.Model(entities.Product{}).Where("id = ?", p.ID).Update("tags = ?", string(tagsJSON))
	log.Printf("product %s with id %d is added\n", p.Name, p.ID)
	pr.prdb.Create(&entities.Activity{
		Type:    entities.AddProductActivity,
		Payload: fmt.Sprintf("محصول %s با شناسه %d اضافه شد", p.Name, p.ID),
	})
	return nil
}

func (pr *ProductRep) GetProduct(ID uint) (entities.Product, error) {
	var product entities.Product
	res := pr.prdb.Preload("Category").First(&product, ID)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return product, fmt.Errorf("couldn't find product")
		}
		return product, fmt.Errorf("error occured, %v", res.Error)
	}
	return product, nil
}

func (pr *ProductRep) GetProducts() ([]entities.Product, error) {
	var products []entities.Product
	tx := pr.prdb.Find(&products)
	if tx.Error != nil {
		return products, fmt.Errorf("couldn't get products")
	}
	return products, nil
}

func (pr *ProductRep) EditProduct(p entities.Product) error {
	tx := pr.prdb.Where("id = ?", p.ID).Updates(entities.Product{
		Name:        p.Name,
		Image:       p.Image,
		Images:      p.Images,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		Info:        p.Info,
		Off:         p.Off,
		CategoryID:  p.CategoryID,
		Tags:        p.Tags,
	},
	)
	if tx.Error != nil {
		return fmt.Errorf("couldn't update product")
	}
	return nil
}

func (pr *ProductRep) DeleteProduct(ID uint) error {
	var product entities.Product
	tx := pr.prdb.First(&product, ID)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("product not found")
	}
	tx = pr.prdb.Delete(&product)
	if tx.Error != nil {
		return fmt.Errorf("couldn't delete product")
	}
	return nil
}
