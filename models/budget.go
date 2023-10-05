package models

import (
	"time"

	"github.com/google/uuid"
)

type Budget struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	UserID          uuid.UUID
	ClientID        uint
	BudgetOfferings []BudgetOffering
	Price           float32
	Date            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type CreateBudgetRequest struct {
	UserID   uuid.UUID `json:"userID"`
	ClientID uint      `json:"clientID"`
	Price    float32   `json:"price"`
}
