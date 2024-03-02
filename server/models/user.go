package models

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string     `gorm:"not null" json:"name"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	Customers []Customer `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"customers"`
	Jobs      []Job      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"jobs"`
	Budgets   []Budget   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"budgets"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `json:"id"`
	Name     *string   `json:"name,omitempty"`
	Password *string   `json:"password,omitempty"`
}

type LoginUserResquest struct {
	Email    string
	Password string
}

type ListUserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
}