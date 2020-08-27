package entity

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	BaseEntity
	ID          string   `gorm:"column:id;primary_key"`
	ProductName string   `gorm:"column:product_name;not null"`
	ProductCode string   `gorm:"column:product_code;not null"`
	Category    Category `gorm:"foreignkey:CategoryId"`
	CategoryId  string
}

func (p *Product) TableName() string {
	return "m_product"
}
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = guuid.New().String()
	return nil
}
func (p *Product) ToString() string {
	product, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(product)
}

type ProductWithCategoryDTO struct {
	Id           string
	ProductCode  string
	ProductName  string
	CategoryId   string
	CategoryName string
}
