package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
}

type LoginUserResquest struct {
	Email    string
	Password string
}
