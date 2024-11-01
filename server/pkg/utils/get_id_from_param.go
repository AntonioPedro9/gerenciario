package utils

import (
	"server/pkg/errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIdFromParam(c *gin.Context) (uint, error) {
	paramId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, errors.InvalidParamIdError
	}
	return uint(paramId), nil
}
