package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessTokenString, err := c.Cookie("Authorization")
		if err != nil {
			return
		}

		accessToken, err := jwt.Parse(accessTokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})
		if err != nil || !accessToken.Valid {
			if err == err.(*jwt.ValidationError) {
				// implementar um novo access token a partir do refresh token
			}
			return
		}

		claims, ok := accessToken.Claims.(jwt.MapClaims)
		if !ok {
			return
		}

		userId := uint(claims["sub"].(float64))

		c.Set("x-user-id", userId)
		c.Next()
	}
}
