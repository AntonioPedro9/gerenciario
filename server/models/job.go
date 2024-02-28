package models

import "github.com/google/uuid"

type Job struct {
	UserID      uuid.UUID `gorm:"not null" json:"userID"`
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	Price       float32   `json:"price"`
}

type CreateJobRequest struct {
	UserID      uuid.UUID `json:"userID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    uint      `json:"duration"`
	Price       float32   `json:"price"`
}

type UpdateJobRequest struct {
	UserID      uuid.UUID `json:"userID"`
	ID          uint      `json:"id"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Duration    *uint     `json:"duration,omitempty"`
	Price       *float32  `json:"price,omitempty"`
}
