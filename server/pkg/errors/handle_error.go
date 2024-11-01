package errors

import (
	"fmt"
	"net/http"
	"server/pkg/logs"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleError(c *gin.Context, err error) {
	if customError, ok := err.(*CustomError); ok {
		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
	} else if validatorError, ok := err.(validator.ValidationErrors); ok {
		validationErrors := formatValidationErrors(validatorError)
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
	} else {
		logs.LogError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro do servidor interno"})
	}
}

func formatValidationErrors(ve validator.ValidationErrors) map[string]string {
	errors := make(map[string]string)
	for _, err := range ve {
		var message string
		switch err.Tag() {
		case "required":
			message = fmt.Sprintf("%s é obrigatório", strings.ToLower(err.Field()))
		case "min":
			message = fmt.Sprintf("%s deve ter pelo menos %s caracteres", strings.ToLower(err.Field()), err.Param())
		case "max":
			message = fmt.Sprintf("%s deve ter no máximo %s caracteres", strings.ToLower(err.Field()), err.Param())
		case "email":
			message = "formato de email inválido"
		default:
			message = fmt.Sprintf("%s é inválido", strings.ToLower(err.Field()))
		}
		errors[strings.ToLower(err.Field())] = message
	}
	return errors
}
