package utils

import (
	"regexp"
	"strings"
)

func FormatEmail(email string) (string, error) {
	if email == "" {
		return email, nil
	}

	email = strings.ToLower(email)

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return "", InvalidEmailError
	}

	return email, nil
}
