package models

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"type:varchar(120);not null"`
	Email    string `gorm:"type:varchar(120);unique;not null"`
	Password string `gorm:"type:varchar(120);not null"`
}

type CreateUserRequest struct {
	Name     string `json:"name"     validate:"required,min=3"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=120"`
}

type GetUserResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserDataRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type UpdateUserPasswordRequest struct {
	Password string `json:"password" validate:"required,min=8,max=120"`
}
