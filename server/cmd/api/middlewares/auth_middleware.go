package middlewares

import (
	"server/internals/repositories"
	"server/pkg/errors"
	"server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userRepository *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId, err := utils.GetUserIdFromToken(c)
		if err != nil {
			errors.HandleError(c, err)
			c.Abort()
			return
		}

		user, err := userRepository.GetById(userId)
		if err != nil {
			errors.HandleError(c, errors.NotFoundError)
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
