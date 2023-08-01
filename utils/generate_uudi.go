package utils

import "github.com/google/uuid"

func GenerateUUDI() string {
	return uuid.NewString()
}
