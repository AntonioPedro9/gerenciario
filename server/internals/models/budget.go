package models

import (
	"time"
)

type Budget struct {
	Id             uint      `gorm:"primaryKey;autoIncrement"`
	UserId         uint      `gorm:"not null"`
	CustomerId     uint      `gorm:"not null"`
	TotalPrice     float32   `gorm:"type:numeric(10,2);not null"`
	Discount       float32   `gorm:"type:numeric(10,2)"`
	Date           time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ExpirationDate time.Time `gorm:""`

	User     User     `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Customer Customer `gorm:"foreignKey:CustomerId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Jobs     []Job    `gorm:"many2many:budget_jobs;constraint:OnDelete:CASCADE"`
}

type CreateBudgetRequest struct {
	CustomerId     uint      `json:"customer_id"     validate:"required"`
	TotalPrice     float32   `json:"total_price"     validate:"required,gt=0"`
	Discount       float32   `json:"discount"        validate:"gte=0"`
	Date           time.Time `json:"date"            validate:"omitempty,gtfield=Date"`
	ExpirationDate time.Time `json:"expiration_date" validate:"omitempty,gtfield=Date"`
	Jobs           []Job     `json:"jobs"`
}

type GetBudgetResponse struct {
	Id             uint      `json:"id"`
	UserId         uint      `json:"user_id"`
	CustomerId     uint      `json:"customer_id"`
	TotalPrice     float32   `json:"total_price"`
	Discount       float32   `json:"discount"`
	Date           time.Time `json:"date"`
	ExpirationDate time.Time `json:"expiration_date"`
	Jobs           []Job     `json:"jobs"`
}

type GetUserBudgetsResponse struct {
	Budgets []GetBudgetResponse `json:"budgets"`
}
