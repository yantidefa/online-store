package models

import "github.com/google/uuid"

type Payment struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	ProductId uuid.UUID `json:"product_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId    uuid.UUID `json:"user_id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Qty       int       `json:"qty"`
	Price     int       `json:"price"`
	CreatedAt *string   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *string   `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *string   `json:"deleted_at" gorm:"column:deleted_at"`
}
type ProductPayment struct {
	ProductName  string `json:"product_name"`
	CategoryName string `json:"category_name"`
	Qty          int    `json:"qty"`
	Price        int    `json:"price"`
}
type UserPayment struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type Checkout struct {
	User       UserPayment      `json:"user"`
	Product    []ProductPayment `json:"product"`
	TotalPrice int              `json:"total_price"`
}
