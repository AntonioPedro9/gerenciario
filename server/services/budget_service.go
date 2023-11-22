package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type BudgetService struct {
	budgetRepository *repositories.BudgetRepository
}

func NewBudgetService(budgetRepository *repositories.BudgetRepository) *BudgetService {
	return &BudgetService{budgetRepository}
}

func (bs *BudgetService) CreateBudget(budget *models.CreateBudgetRequest) error {
	validBudget := &models.Budget{
		Price:    budget.Price,
		UserID:   budget.UserID,
		ClientID: budget.ClientID,
	}

	err := bs.budgetRepository.Create(validBudget)
	if err != nil {
		return err
	}

	for _, serviceID := range budget.ServiceIDs {
		budgetService := models.BudgetService{
			BudgetID:  validBudget.ID,
			ServiceID: serviceID,
		}

		err = bs.budgetRepository.CreateBudgetService(&budgetService)
		if err != nil {
			return err
		}
	}

	return nil
}

func (bs *BudgetService) ListBudgets(userID uuid.UUID) ([]models.Budget, error) {
	return bs.budgetRepository.List(userID)
}

func (bs *BudgetService) DeleteBudget(budgetID uint, tokenID uuid.UUID) error {
	existingBudget, err := bs.budgetRepository.GetBudgetById(budgetID)
	if err != nil {
		return err
	}

	if existingBudget.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	return bs.budgetRepository.Delete(budgetID)
}
