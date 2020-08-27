package entity

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	BaseEntity
	ID           string    `gorm:"column:id;primary_key"`
	CategoryName string    `gorm:"column:category_name;not null"`
	Products     []Product `gorm:"foreignKey:CategoryId"`
}

func (c *Category) TableName() string {
	return "m_category"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = guuid.New().String()
	return nil
}
func (c *Category) ToString() string {
	category, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(category)
}
