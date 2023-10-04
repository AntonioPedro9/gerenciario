package utils

import (
	"errors"
	"regexp"
)

func FormatCPF(cpf string) (string, error) {
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(cpf, "")

	if len(digits) != 11 {
		return "", errors.New("Invalid CPF")
	}

	return digits, nil
}
