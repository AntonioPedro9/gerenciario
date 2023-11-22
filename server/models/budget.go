package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Price          float32
	Date           time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UserID         uuid.UUID `gorm:"not null"`
	ClientID       uint      `gorm:"not null"`
	BudgetServices []BudgetService
}

type CreateBudgetRequest struct {
	Price      float32   `json:"price"`
	UserID     uuid.UUID `json:"userID"`
	ClientID   uint      `json:"clientID"`
	ServiceIDs []uint    `json:"serviceIDs"`
}

type BudgetService struct {
	BudgetID  uint
	ServiceID uint
}
