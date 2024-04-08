package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
    UserID        uuid.UUID   `gorm:"not null" json:"userID"`
    CustomerID    uint        `gorm:"not null" json:"customerID"`
    ID            uint        `gorm:"primaryKey;autoIncrement" json:"id"`
    BudgetJobs    []BudgetJob `gorm:"foreignKey:BudgetID;constraint:OnDelete:CASCADE;" json:"budgetJobs"`
    BudgetDate    time.Time   `gorm:"type:date;default:CURRENT_TIMESTAMP" json:"budgetDate"`
    ScheduledDate time.Time   `gorm:"type:date" json:"scheduledDate"`
    ScheduledTime time.Time   `gorm:"type:time" json:"scheduledTime"`
    Vehicle       string      `gorm:"not null" json:"vehicle"`
    LicensePlate  string      `gorm:"not null" json:"licensePlate"`
    Price         float32     `json:"price"`
    Completed     bool        `gorm:"default:false" json:"completed"`
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
	ScheduledTime time.Time   `json:"scheduledTime"`
	Vehicle       string      `json:"vehicle"`
	LicensePlate  string      `json:"licensePlate"`
	Price         float32     `json:"price"`
	Completed     bool        `json:"completed"`
}

type BudgetJob struct {
	BudgetID uint
	JobID    uint
}
