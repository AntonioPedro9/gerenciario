package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
