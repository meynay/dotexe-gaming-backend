package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

//	type Category struct {
//		ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
//		Name     string             `json:"name" bson:"name"`
//		Image    string             `json:"image" bson:"image"`
//		ParentID primitive.ObjectID `json:"parent_id" bson:"parent_id"`
//	}
type Category struct {
	gorm.Model
	Name         string     `gorm:"type:varchar(50);unique;not null" json:"name"`
	Image        string     `gorm:"type:text" json:"image"`
	ParentID     *uint      `gorm:"index;default:null" json:"parent_id,omitempty"`
	Parent       *Category  `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parent,omitempty"`
	Children     []Category `gorm:"foreignKey:ParentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"children,omitempty"`
	ProductCount int        `gorm:"-:all" json:"product_count"`
}

//	type Product struct {
//		ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
//		Name          string             `json:"name" bson:"name"`
//		Image         string             `json:"image" bson:"image"`
//		Images        []string           `json:"images" bson:"images"`
//		Description   string             `json:"description" bson:"description"`
//		Price         int                `json:"price" bson:"price"`
//		Stock         int                `json:"stock" bson:"stock"`
//		Off           float64            `json:"off" bson:"off"`
//		Info          map[string]string  `json:"info" bson:"info"`
//		CategoryID    primitive.ObjectID `json:"category_id" bson:"category_id"`
//		AddedAt       time.Time          `json:"time_added" bson:"time_added"`
//		UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at"`
//		Rating        float64            `json:"rating" bson:"rating"`
//		RateCount     int                `json:"rate_count" bson:"rate_count"`
//		Views         int                `json:"views" bson:"views"`
//		PurchaseCount int                `json:"purchase_count" bson:"purchase_count"`
//		Tags          []string           `json:"tags" bson:"tags"`
//	}
type JSONB map[string]interface{} // or map[string]string

func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan JSONB: %v", value)
	}
	return json.Unmarshal(bytes, j)
}

func (j JSONB) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

type Product struct {
	gorm.Model
	Name          string   `gorm:"type:varchar(100);not null;index" json:"name"`
	Image         string   `gorm:"type:text" json:"image"`
	Images        JSONB    `gorm:"type:jsonb" json:"images"`
	Description   string   `gorm:"type:text;not null" json:"description"`
	Price         int      `gorm:"type:integer;not null" json:"price"`
	Stock         int      `gorm:"type:integer;not null;default:0" json:"stock"`
	Off           float64  `gorm:"type:numeric(5,2);default:0.0" json:"off"`
	Info          JSONB    `gorm:"type:jsonb" json:"info"`
	CategoryID    uint     `gorm:"index;not null" json:"category_id"`
	Category      Category `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category,omitempty"`
	Rating        float64  `gorm:"type:numeric(3,2);default:0.0" json:"rating"`
	RateCount     int      `gorm:"type:integer;default:0" json:"rate_count"`
	Views         int      `gorm:"type:integer;default:0" json:"views"`
	PurchaseCount int      `gorm:"type:integer;default:0" json:"purchase_count"`
	Tags          JSONB    `gorm:"type:jsonb" json:"tags"`
}

type ProductLess struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Image         string  `json:"image"`
	Description   string  `json:"description"`
	Price         int     `json:"price"`
	Category      string  `json:"category"`
	Off           float64 `json:"off"`
	Rating        float64 `json:"rating"`
	RateCount     int     `json:"rate_count"`
	Views         int     `json:"views"`
	PurchaseCount int     `json:"purchase_count"`
}

type Filter struct {
	Query         string `json:"query"`
	CategoryID    uint   `json:"category_id"`
	Page          int    `json:"page"`
	NumberOfItems int    `json:"number_of_items"`
	Order         int    `json:"order"`
	Available     bool   `json:"only_available"`
}

type PScore struct {
	Pr    Product
	Score float64
}

const (
	CheapToExpensive = iota
	ExpensiveToCheap
	MostOffToLeast
	Newest
	MostViewed
	MostPurchased
	MostRelevant
	MostRate
)
