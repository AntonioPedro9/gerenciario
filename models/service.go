package models

import "github.com/google/uuid"

type Service struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	Description string
	Duration    uint
	Price       float32
	UserID      uuid.UUID `gorm:"not null"`
}

type CreateServiceRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	Price       float32   `json:"price"`
	UserID      uuid.UUID `json:"userID"`
}

type UpdateServiceRequest struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	Price       float32   `json:"price"`
	UserID      uuid.UUID `json:"userID"`
}
