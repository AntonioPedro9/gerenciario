package utils

import (
	"os"
	"server/pkg/errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(sub uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.TokenGenerationError
	}

	return tokenString, nil
}
