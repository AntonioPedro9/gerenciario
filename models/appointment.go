package models

import (
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	Date     time.Time `gorm:"not null"`
	BudgetID uint      `gorm:"not null"`
	UserID   uuid.UUID `gorm:"not null"`
}

type CreateAppointmentRequest struct {
	Date     time.Time `json:"date"`
	BudgetID uint      `json:"budgetID"`
	UserID   uuid.UUID `json:"userID"`
}

type UpdateAppointmentRequest struct {
	ID     uint      `json:"id"`
	Date   *time.Time `json:"date,omitempty"`
	UserID uuid.UUID `json:"userID"`
}
