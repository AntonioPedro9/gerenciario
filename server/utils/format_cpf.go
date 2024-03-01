package utils

import (
	"regexp"
)

func FormatCPF(cpf string) (string, error) {
	re := regexp.MustCompile(`\D`)
	digits := re.ReplaceAllString(cpf, "")

	if len(digits) != 11 {
		return "", InvalidCpfError
	}

	return digits, nil
}
