package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"unique"`
	Password string    `gorm:"not null"`
	Clients  []Client
	Services []Service
	Budgets  []Budget
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     *string    `json:"name,omitempty"`
	Password *string    `json:"password,omitempty"`
}

type LoginUserResquest struct {
	Email    string
	Password string
}
