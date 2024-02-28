package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	UserID        uuid.UUID   `gorm:"not null" json:"userID"`
	CustomerID    uint        `gorm:"not null" json:"customerID"`
	ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	BudgetJobs    []BudgetJob `gorm:"constraint:OnDelete:CASCADE;" json:"budgetJobs"`
	BudgetDate    time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"budgetDate"`
	ScheduledDate time.Time   `json:"schedeledDate"`
	Vehicle       string      `gorm:"not null" json:"vehicle"`
	LicensePlate  string      `gorm:"not null" json:"licensePlate"`
	Price         float32     `json:"price"`
}

type CreateBudgetRequest struct {
	UserID       uuid.UUID `json:"userID"`
	CustomerID   uint      `json:"customerID"`
	JobIDs       []uint    `json:"jobIDs"`
	Vehicle      string    `json:"vehicle"`
	LicensePlate string    `json:"licensePlate"`
	Price        float32   `json:"price"`
}

type ListBudgetsResponse struct {
	UserID        uuid.UUID   `json:"userID"`
	ID            uint        `json:"id"`
	CustomerID    uint        `json:"customerID"`
	CustomerName  string      `json:"customerName"`
	CustomerPhone string      `json:"customerPhone"`
	BudgetJobs    []BudgetJob `json:"budgetJobs"`
	BudgetDate    time.Time   `json:"budgetDate"`
	ScheduledDate time.Time   `json:"schedeledDate"`
	Vehicle       string      `json:"vehicle"`
	LicensePlate  string      `json:"licensePlate"`
	Price         float32     `json:"price"`
}

type BudgetJob struct {
	BudgetID uint
	JobID    uint
}
