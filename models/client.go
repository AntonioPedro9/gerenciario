package models

import "github.com/google/uuid"

type Client struct {
	ID     uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CPF    string    `gorm:"unique" json:"cpf"`
	Name   string    `gorm:"not null" json:"name"`
	Email  string    `json:"email"`
	Phone  string    `gorm:"not null" json:"phone"`
	UserID uuid.UUID `gorm:"not null" json:"userID"`
}

type CreateClientRequest struct {
	CPF    string    `json:"cpf"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`
	UserID uuid.UUID `json:"userID"`
}

type UpdateClientRequest struct {
	ID     uint       `json:"id"`
	CPF    *string    `json:"cpf,omitempty"`
	Name   *string    `json:"name,omitempty"`
	Email  *string    `json:"email,omitempty"`
	Phone  *string    `json:"phone,omitempty"`
	UserID uuid.UUID  `json:"userID"`
}
