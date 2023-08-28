package models

import (
	"server/utils"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUser(name, email, password string) *User {
	userId := utils.GenerateUUDI()
	hashedPassword := utils.HashPassword(password)

	return &User{
		ID:       userId,
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}
}
