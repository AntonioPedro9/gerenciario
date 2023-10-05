package models

import "github.com/google/uuid"

type Service struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	Duration    uint
	UserID      uuid.UUID
}

type CreateServiceRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	UserID      uuid.UUID `json:"userID"`
}

type UpdateServiceRequest struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	UserID      uuid.UUID `json:"userID"`
}
