package utils

import (
	"regexp"
	"strings"
)

func IsValidName(name string) bool {
	name = strings.Trim(name, " ")
	nameRegex := `^[A-Za-z0-9 ]{3,}$`
	match, _ := regexp.MatchString(nameRegex, name)

	return match
}
