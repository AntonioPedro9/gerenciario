package models

import (
	"server/utils"
	"time"
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
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, password string) *User {
	userId := utils.GenerateUUDI()
	createdAt := utils.GetCurrentTime()

	return &User{
		ID:        userId,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
	}
}
