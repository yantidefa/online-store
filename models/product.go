package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CategoryId  uuid.UUID `json:"category_id" gorm:"column:category_id"`
	Name        string    `json:"name" gorm:"column:name"`
	Description string    `json:"description" gorm:"column:description"`
	Stock       int       `json:"stock" gorm:"column:stock"`
	Price       float64   `json:"price" gorm:"column:price"`
	CreatedAt   *string   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *string   `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *string   `json:"deleted_at" gorm:"column:deleted_at"`
}
