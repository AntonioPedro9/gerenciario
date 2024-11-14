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
	Refresh(c *gin.Context)
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

	accessToken, refreshToken, err := ah.authService.Auth(&user)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	cookieExpirationTime := 3600 * 24 * 30
	c.SetCookie("Authorization", accessToken, cookieExpirationTime, "", "", false, true)
	c.SetCookie("RefreshToken", refreshToken, cookieExpirationTime, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken, 
		"refreshToken": refreshToken,
	})
}

func (ah *AuthHandler) Refresh(c *gin.Context) {
	refreshTokenString, err := c.Cookie("RefreshToken")
	if err != nil {
		errors.HandleError(c, errors.CookieNotFoundError)
		return
	}

	newAccessToken, err := ah.authService.Refresh(refreshTokenString)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	cookieExpirationTime := 3600 * 24 * 30
	c.SetCookie("Authorization", newAccessToken, cookieExpirationTime, "", "", false, true)
	c.SetSameSite(http.SameSiteLaxMode)
	c.JSON(http.StatusOK, gin.H{"token": newAccessToken})
}
