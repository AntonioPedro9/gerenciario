package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.NewString()
	return id
}
