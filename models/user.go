package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Token     string    `json:"token"`
	IsLogin   bool      `json:"is_login"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt *string   `json:"created_at"`
	UpdatedAt *string   `json:"updated_at"`
	DeletedAt *string   `json:"deleted_at"`
}
type Login struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}
