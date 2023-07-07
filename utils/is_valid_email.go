package utils

import (
	"regexp"
	"strings"
)

func IsValidEmail(email string) bool {
	email = strings.Trim(email, " ")
	emailRegex := `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)

	return match
}
