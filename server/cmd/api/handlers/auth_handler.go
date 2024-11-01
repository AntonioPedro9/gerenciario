package handlers

import (
	"net/http"
	"server/internals/models"
	services "server/internals/services"
	"server/pkg/errors"

	"github.com/gin-gonic/gin"
)

type AuthHandlerInterface interface {
	Auth(c *gin.Context)
}

type AuthHandler struct {
	authService services.AuthServiceInterface
}

func NewAuthHandler(authService services.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService}
}

func (ah *AuthHandler) Auth(c *gin.Context) {
	var user models.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	tokenString, err := ah.authService.Auth(&user)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	expirationTime := 3600 * 24 * 30
	c.SetCookie("Authorization", tokenString, expirationTime, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
