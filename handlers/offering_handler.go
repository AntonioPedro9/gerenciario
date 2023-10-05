package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type OfferingHandler struct {
	offeringService *services.OfferingService
}

func NewOfferingHandler(offeringService *services.OfferingService) *OfferingHandler {
	return &OfferingHandler{offeringService}
}

func (oh *OfferingHandler) CreateOffering(c *gin.Context) {
	var offering models.CreateOfferingRequest
	if err := c.ShouldBindJSON(&offering); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := oh.offeringService.CreateOffering(&offering); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (oh *OfferingHandler) ListOfferings(c *gin.Context) {
	id := c.Param("userID")

	userID, err := uuid.Parse(id)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	authUserID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	offerings, err := oh.offeringService.ListOfferings(userID, authUserID)
	if err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusOK, offerings)
}

func (oh *OfferingHandler) UpdateOffering(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	authUserID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var offering models.UpdateOfferingRequest
	if err := c.ShouldBindJSON(&offering); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := oh.offeringService.UpdateOffering(&offering, authUserID); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (oh *OfferingHandler) DeleteOffering(c *gin.Context) {
	id := c.Param("offeringID")

	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offering ID"})
		return
	}
	offeringID := uint(parsedID)

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	authUserID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := oh.offeringService.DeleteOffering(offeringID, authUserID); err != nil {
		log.Error(err)

		customError, ok := err.(*utils.CustomError)
		if !ok {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(customError.StatusCode, gin.H{"error": customError.Message})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
