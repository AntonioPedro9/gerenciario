package models

type User struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Email    string `gorm:"unique"`
	Password string
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUserResquest struct {
	Email    string
	Password string
}
