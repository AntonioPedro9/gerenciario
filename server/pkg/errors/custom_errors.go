package errors

import (
	"net/http"
)

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
}

var (
	// 400 errors
	JsonBindingError    = NewCustomError(http.StatusBadRequest, "dados enviados inválidos")
	InvalidParamIdError = NewCustomError(http.StatusBadRequest, "ID do parâmetro inválido")

	// 401 errors
	UnauthorizedError      = NewCustomError(http.StatusUnauthorized, "ação não autorizada")
	IncorrectPasswordError = NewCustomError(http.StatusUnauthorized, "senha incorreta")
	CookieNotFoundError    = NewCustomError(http.StatusUnauthorized, "token de autorização não encontrado")
	TokenParsingError      = NewCustomError(http.StatusUnauthorized, "falha ao processar o token de autorização")
	InvalidTokenError      = NewCustomError(http.StatusUnauthorized, "token de autorização inválido")
	ExpiredTokenError      = NewCustomError(http.StatusUnauthorized, "token de autorização expirado")

	// 404 errors
	NotFoundError = NewCustomError(http.StatusNotFound, "recurso não encontrado")

	// 409 errors
	EmailInUseError = NewCustomError(http.StatusConflict, "usuário já cadastrado")

	// 500 errors
	TokenGenerationError = NewCustomError(http.StatusInternalServerError, "erro ao gerar o token de autenticação")
	PasswordHashingError = NewCustomError(http.StatusInternalServerError, "erro ao processar a senha")
)
