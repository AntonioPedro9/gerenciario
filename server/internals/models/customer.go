package models

type Customer struct {
	Id     uint   `gorm:"primaryKey;autoIncrement"`
	UserId uint   `gorm:"not null"`
	Name   string `gorm:"type:varchar(120);not null"`
	CPF    string `gorm:"type:varchar(11);unique"`
	Phone  string `gorm:"type:varchar(13);not null"`
	Email  string `gorm:"type:varchar(120)"`

	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CreateCustomerRequest struct {
	Name  string `json:"name"  validate:"required,min=3"`
	CPF   string `json:"cpf"   validate:"omitempty,numeric,len=11"`
	Phone string `json:"phone" validate:"required,numeric,len=13"`
	Email string `json:"email" validate:"omitempty,email"`
}

type GetCustomerResponse struct {
	Id     uint   `json:"id"`
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
}

type GetUserCustomersResponse struct {
	Customers []GetCustomerResponse `json:"customers"`
}

type UpdateCustomerRequest struct {
	Name  string `json:"name"  validate:"required,min=3"`
	CPF   string `json:"cpf"   validate:"omitempty,numeric,len=11"`
	Phone string `json:"phone" validate:"required,numeric,len=13"`
	Email string `json:"email" validate:"omitempty,email"`
}
