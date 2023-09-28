package middlewares

import (
	"net/http"
	"server/database"
	"server/repositories"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	tokenID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	db := database.ConnectToDatabase()
	userRepository := repositories.NewUserRepository(db)

	user, err := userRepository.GetUserById(tokenID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Set("user", user)
	c.Next()
}
