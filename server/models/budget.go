package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID             uint            `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         uuid.UUID       `gorm:"not null" json:"userID"`
	ClientID       uint            `gorm:"not null" json:"clientID"`
	ClientName     string          `gorm:"not null" json:"clientName"`
	ClientPhone    string          `gorm:"not null" json:"clientPhone"`
	Date           time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
	BudgetServices []BudgetService `gorm:"constraint:OnDelete:CASCADE;" json:"budgetServices"`
	Vehicle        string          `gorm:"not null" json:"vehicle"`
	LicensePlate   string          `gorm:"not null" json:"licensePlate"`
	Price          float32         `json:"price"`
}

type CreateBudgetRequest struct {
	UserID       uuid.UUID `json:"userID"`
	ClientID     uint      `json:"clientID"`
	ClientName   string    `json:"clientName"`
	ClientPhone  string    `json:"clientPhone"`
	ServiceIDs   []uint    `json:"serviceIDs"`
	Vehicle      string    `json:"vehicle"`
	LicensePlate string    `json:"licensePlate"`
	Price        float32   `json:"price"`
}

type BudgetService struct {
	BudgetID  uint
	ServiceID uint
}
