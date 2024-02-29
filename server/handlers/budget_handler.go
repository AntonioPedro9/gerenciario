package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type BudgetHandler struct {
	budgetService *services.BudgetService
}

func NewBudgetHandler(budgetService *services.BudgetService) *BudgetHandler {
	return &BudgetHandler{budgetService}
}

func (bh *BudgetHandler) CreateBudget(c *gin.Context) {
	var budget models.CreateBudgetRequest
	if err := c.ShouldBindJSON(&budget); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON request"})
		return
	}

	if err := bh.budgetService.CreateBudget(&budget); err != nil {
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

func (bh *BudgetHandler) ListBudgets(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	userID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	budgets, err := bh.budgetService.ListBudgets(userID)
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

	c.JSON(http.StatusOK, budgets)
}

func (bh *BudgetHandler) GetBudget(c *gin.Context) {
	paramBudgetID := c.Param("budgetID")

	parsedBudgetID, err := strconv.ParseUint(paramBudgetID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid budget ID"})
		return
	}
	budgetID := uint(parsedBudgetID)

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	tokenID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	budget, err := bh.budgetService.GetBudget(budgetID, tokenID)
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

	c.JSON(http.StatusOK, budget)
}

func (bh *BudgetHandler) GetBudgetServices(c *gin.Context) {
	paramBudgetID := c.Param("budgetID")

	parsedBudgetID, err := strconv.ParseUint(paramBudgetID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid budaget ID"})
		return
	}
	budgetID := uint(parsedBudgetID)

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	tokenID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	budgetServices, err := bh.budgetService.GetBudgetServices(budgetID, tokenID)
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

	c.JSON(http.StatusOK, budgetServices)
}

func (bh *BudgetHandler) DeleteBudget(c *gin.Context) {
	paramBudgetID := c.Param("budgetID")

	parsedID, err := strconv.ParseUint(paramBudgetID, 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid budget ID"})
		return
	}
	budgetID := uint(parsedID)

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
		return
	}

	tokenID, err := utils.GetIDFromToken(tokenString)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if err := bh.budgetService.DeleteBudget(budgetID, tokenID); err != nil {
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
