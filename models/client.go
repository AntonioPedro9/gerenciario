package models

import "github.com/google/uuid"

type Client struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Name   string
	Email  string
	Phone  string
	UserID uuid.UUID
}

type CreateClientRequest struct {
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`
	UserID uuid.UUID `json:"userID"`
}

type UpdateClientRequest struct {
	ID     uint      `json:"id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`
	UserID uuid.UUID `json:"userID"`
}
