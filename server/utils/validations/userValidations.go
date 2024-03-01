package validations

import (
	"regexp"
	"server/models"
	"server/utils"
)

func ValidateCreateUserRequest(user *models.CreateUserRequest) error {
	if len(user.Name) < 2 {
		return utils.InvalidNameError
	}

	if len(user.Email) == 0 {
		return utils.InvalidEmailError
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		return utils.InvalidEmailError
	}

	if len(user.Password) < 8 || len(user.Password) > 128 {
		return utils.PasswordLengthError
	}

	return nil
}
