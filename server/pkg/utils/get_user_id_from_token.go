package utils

import (
	"os"
	"server/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GetUserIdFromToken(c *gin.Context) (uint, error) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		return 0, errors.CookieNotFoundError
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, errors.TokenParsingError
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.InvalidTokenError
	}

	expirationTime := int64(claims["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		return 0, errors.ExpiredTokenError
	}

	userId := uint(claims["sub"].(float64))

	return userId, nil
}
