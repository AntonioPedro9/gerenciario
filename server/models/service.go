package models

import "github.com/google/uuid"

type Service struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	Price       float32   `json:"price"`
	UserID      uuid.UUID `gorm:"not null" json:"userID"`
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
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Duration    *uint     `json:"duration,omitempty"`
	Price       *float32  `json:"price,omitempty"`
	UserID      uuid.UUID `json:"userID"`
}
