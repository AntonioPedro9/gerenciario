package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"server/database"
	"server/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	expirationTime := int64(claims["exp"].(float64))
	if time.Now().Unix() > expirationTime {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID, ok := claims["sub"].(string)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userRepository := repositories.NewUserRepository(database.DB)

	user, err := userRepository.GetUserById(userID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("user", user)
	c.Next()
}
