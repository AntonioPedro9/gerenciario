package utils

import (
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func GenerateUUDI() string {
	id, err := uuid.NewRandom()
	if err != nil {
		log.Error("Failed to generate UUID: ", err)
		return ""
	}

	return id.String()
}
