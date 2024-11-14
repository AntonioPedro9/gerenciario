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
	switch e := err.(type) {
	case *CustomError:
		handleCustomError(c, e)
	case validator.ValidationErrors:
		handleValidationError(c, e)
	default:
		handleServerError(c, err)
	}
}

func handleCustomError(c *gin.Context, err *CustomError) {
	c.JSON(err.StatusCode, gin.H{"error": err.Message})
}

func handleValidationError(c *gin.Context, err validator.ValidationErrors) {
	validationErrors := formatValidationErrors(err)
	c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
}

func handleServerError(c *gin.Context, err error) {
	logs.LogError(err)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "erro do servidor interno"})
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
