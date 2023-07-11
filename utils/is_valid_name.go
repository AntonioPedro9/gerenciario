package utils

import (
	"strings"
	"unicode"
)

func IsValidName(name string) bool {
	name = strings.TrimSpace(name)

	if name == "" {
		return false
	}

	if len(name) < 2 || len(name) > 64 {
		return false
	}

	for _, char := range name {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false
		}
	}

	return true
}
