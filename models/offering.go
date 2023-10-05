package models

import "github.com/google/uuid"

type Offering struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	Description string
	Duration    uint
	UserID      uuid.UUID
}

type CreateOfferingRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	UserID      uuid.UUID `json:"userID"`
}

type UpdateOfferingRequest struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	UserID      uuid.UUID `json:"userID"`
}
