package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamID(paramName string, c *gin.Context) (uint, error) {
	paramID := c.Param(paramName)

	jobID, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint(jobID), nil
}
