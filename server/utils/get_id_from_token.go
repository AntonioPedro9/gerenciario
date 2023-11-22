package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GetIDFromToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return uuid.Nil, errors.New("Invalid token")
	}

	expirationTime := int64(claims["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		return uuid.Nil, errors.New("Expired token")
	}

	tokenID, ok := claims["sub"].(string)
	if !ok {
		return uuid.Nil, errors.New("No 'sub' claim found")
	}

	parsedTokenID, err := uuid.Parse(tokenID)
	if err != nil {
		return uuid.Nil, err
	}

	return parsedTokenID, nil
}
