package models

import "github.com/google/uuid"

type CategoryProduct struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CategoryName string    `json:"category_name" gorm:"column:name"`
	CreatedAt    *string   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt    *string   `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt    *string   `json:"deleted_at" gorm:"column:deleted_at"`
}
