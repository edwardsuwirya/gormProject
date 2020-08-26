package entity

import (
	"encoding/json"
	guuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Category struct {
	BaseEntity
	ID string `gorm:"column:id;primary_key"`
	CategoryName string `gorm:"column:category_name;not null"`
}
func (c *Category) TableName() string {
	return "m_category"
}

func (c *Category) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", guuid.New().String())
	return nil
}
func (c *Category) ToString() string {
	category, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(category)
}
