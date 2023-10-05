package models

import (
	"time"

	"github.com/google/uuid"
)

type Contract struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	UserID           uuid.UUID
	ClientID         uint
	ContractServices []ContractService
	Price            float32
	Date             time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type CreateContractRequest struct {
	UserID    uuid.UUID `json:"userID"`
	ClientID  uint      `json:"clientID"`
	Price     float32   `json:"price"`
}
