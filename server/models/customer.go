package models

import "github.com/google/uuid"

type Customer struct {
	UserID  uuid.UUID        `gorm:"not null" json:"userID"`
	ID      uint             `gorm:"primaryKey;autoIncrement" json:"id"`
	CPF     string           `json:"cpf"`
	Name    string           `gorm:"not null" json:"name"`
	Email   string           `json:"email"`
	Phone   string           `gorm:"not null" json:"phone"`
	Budgets []CustomerBudget `gorm:"constraint:OnDelete:CASCADE;" json:"budgets"`
}

type CreateCustomerRequest struct {
	UserID uuid.UUID `json:"userID"`
	CPF    string    `json:"cpf"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
	Phone  string    `json:"phone"`
}

type UpdateCustomerRequest struct {
	UserID uuid.UUID  `json:"userID"`
	ID     uint       `json:"id"`
	CPF    *string    `json:"cpf,omitempty"`
	Name   *string    `json:"name,omitempty"`
	Email  *string    `json:"email,omitempty"`
	Phone  *string    `json:"phone,omitempty"`
}

type CustomerBudget struct {
	CustomerID uint
	BudgetID   uint
}