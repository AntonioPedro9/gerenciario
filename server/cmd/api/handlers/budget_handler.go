package handlers

import (
	"net/http"
	"server/internals/models"
	services "server/internals/services"
	"server/pkg/errors"
	"server/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BudgetHandlerInterface interface {
	CreateBudget(c *gin.Context)
	GetBudget(c *gin.Context)
	GetUserBudgets(c *gin.Context)
	DeleteBudget(c *gin.Context)
}

type BudgetHandler struct {
	budgetService services.BudgetServiceInterface
}

func NewBudgetHandler(budgetService services.BudgetServiceInterface) *BudgetHandler {
	return &BudgetHandler{budgetService}
}

func (bh *BudgetHandler) CreateBudget(c *gin.Context) {
	var data models.CreateBudgetRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errors.HandleError(c, errors.JsonBindingError)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := bh.budgetService.CreateBudget(&data, tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (bh *BudgetHandler) GetBudget(c *gin.Context) {
	budgetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	budget, err := bh.budgetService.GetBudget(uint(budgetId), tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, budget)
}

func (bh *BudgetHandler) GetUserBudgets(c *gin.Context) {
	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	budgets, err := bh.budgetService.GetUserBudgets(tokenUserId)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, budgets)
}

func (bh *BudgetHandler) DeleteBudget(c *gin.Context) {
	budgetId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	tokenUserId, err := utils.GetUserIdFromAccessToken(c)
	if err != nil {
		errors.HandleError(c, err)
		return
	}

	if err := bh.budgetService.DeleteBudget(uint(budgetId), tokenUserId); err != nil {
		errors.HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
