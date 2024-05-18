package models

import "github.com/google/uuid"

type Cart struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	ProductId uuid.UUID `json:"product_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Qty       int       `json:"qty"`
	CreatedAt *string   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *string   `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *string   `json:"deleted_at" gorm:"column:deleted_at"`
}
