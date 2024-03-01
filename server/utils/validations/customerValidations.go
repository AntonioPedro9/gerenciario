package validations

import (
	"regexp"
	"server/models"
	"server/utils"
)

func ValidateCreateCustomerRequest(customer *models.CreateCustomerRequest) error {
	if len(customer.Name) < 2 {
		return utils.InvalidNameError
	}

	if len(customer.Phone) == 0 {
		return utils.InvalidPhoneError
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(customer.Email) {
		return utils.InvalidEmailError
	}

	return nil
}

func ValidateUpdateCustomerRequest(customer *models.UpdateCustomerRequest) error {
	if customer.Name != nil {
		if len(*customer.Name) < 2 {
			return utils.InvalidNameError
		}
	}

	if customer.Email != nil {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(*customer.Email) {
			return utils.InvalidEmailError
		}
	}

	return nil
}