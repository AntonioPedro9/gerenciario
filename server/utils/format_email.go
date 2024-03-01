package utils

import "regexp"

func FormatEmail(email string) (string, error) {
	if email == "" {
		return email, nil
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return "", InvalidEmailError
	}

	return email, nil
}
