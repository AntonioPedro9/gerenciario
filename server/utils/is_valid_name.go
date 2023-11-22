package utils

import (
	"strings"
)

func IsValidName(name string) bool {
	name = strings.TrimSpace(name)

	if name == "" {
		return false
	}

	if len(name) < 2 || len(name) > 64 {
		return false
	}

	return true
}
