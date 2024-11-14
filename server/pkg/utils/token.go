package utils

import (
	"net/http"
	"os"
	"server/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func generateToken(sub uint, secret string, expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.TokenGenerationError
	}

	return tokenString, nil
}

func GenerateAccessAndRefreshToken(sub uint) (string, string, error) {
	accessToken, err := generateToken(sub, os.Getenv("ACCESS_SECRET"), time.Minute*1)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := generateToken(sub, os.Getenv("REFRESH_SECRET"), time.Hour*24*7)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func verifyToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.NewCustomError(http.StatusUnauthorized, err.Error())
	}
	return token, nil
}

func VerifyAccessToken(tokenString string) (*jwt.Token, error) {
	return verifyToken(tokenString, os.Getenv("ACCESS_SECRET"))
}

func VerifyRefreshToken(tokenString string) (*jwt.Token, error) {
	return verifyToken(tokenString, os.Getenv("REFRESH_SECRET"))
}

func GetUserIdFromAccessToken(c *gin.Context) (uint, error) {
	accessTokenString, err := c.Cookie("Authorization")
	if err != nil {
		return 0, errors.CookieNotFoundError
	}

	accessToken, err := VerifyAccessToken(accessTokenString)
	if err != nil {
		return 0, err
	}

	claims, ok := accessToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.InvalidTokenError
	}

	userId := uint(claims["sub"].(float64))

	return userId, nil
}
