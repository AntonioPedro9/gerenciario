package handlers

import (
	"net/http"
	"server/models"
	"server/services"
	"server/utils"

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
	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
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
	budgetID, err := utils.GetParamID("budgetID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	budget, err := bh.budgetService.GetBudget(budgetID, userID)
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
	budgetID, err := utils.GetParamID("budgetID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	budgetServices, err := bh.budgetService.GetBudgetServices(budgetID, userID)
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
	budgetID, err := utils.GetParamID("budgetID", c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param ID"})
		return
	}

	userID, err := utils.GetUserIdFromToken(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized action"})
		return
	}

	if err := bh.budgetService.DeleteBudget(budgetID, userID); err != nil {
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
