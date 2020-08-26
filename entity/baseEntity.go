package entity

import "time"

type BaseEntity struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" sql:"index"`
}
