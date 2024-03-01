package utils

import "net/http"

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

var (
	InvalidNameError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid name",
	}
	InvalidEmailError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid email",
	}
	PasswordLengthError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "InvalidPasswordLength",
	}
	EmailInUseError = &CustomError{
		StatusCode: http.StatusConflict,
		Message:    "Email already in use",
	}
	UnauthorizedActionError = &CustomError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized action",
	}
	NotFoundError = &CustomError{
		StatusCode: http.StatusNotFound,
		Message:    "Not found",
	}
	InvalidEmailOrPasswordError = &CustomError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid email or password",
	}
	InvalidCpfError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid CPF",
	}
	InvalidPhoneError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid phone number",
	}
	InvalidPriceError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid price",
	}
	InvalidDurationError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid duration",
	}
	InvalidDateError = &CustomError{
		StatusCode: http.StatusBadRequest,
		Message: "Invalid date",
	}
)
