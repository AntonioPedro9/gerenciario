package services

import (
	"server/models"
	"server/repositories"
	"server/utils"
	"strings"

	"github.com/google/uuid"
)

type BudgetService struct {
	budgetRepository *repositories.BudgetRepository
}

func NewBudgetService(budgetRepository *repositories.BudgetRepository) *BudgetService {
	return &BudgetService{budgetRepository}
}

func (bs *BudgetService) CreateBudget(budget *models.CreateBudgetRequest, tokenID uuid.UUID) error {
	if budget.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	formattedVehicle, err := utils.FormatName(budget.Vehicle)
	if err != nil {
		return err
	}

	validBudget := &models.Budget{
		UserID:       budget.UserID,
		CustomerID:   budget.CustomerID,
		Vehicle:      formattedVehicle,
		LicensePlate: strings.ToUpper(budget.LicensePlate),
		Price:        budget.Price,
	}

	err = bs.budgetRepository.Create(validBudget)
	if err != nil {
		return err
	}

	for _, jobID := range budget.JobIDs {
		budgetJob := models.BudgetJob{
			BudgetID: validBudget.ID,
			JobID:    jobID,
		}

		err = bs.budgetRepository.CreateBudgetJob(&budgetJob)
		if err != nil {
			return err
		}
	}

	return nil
}


func (bs *BudgetService) ListBudgets(userID uuid.UUID) ([]models.ListBudgetsResponse, error) {
	return bs.budgetRepository.List(userID)
}

func (bs *BudgetService) GetBudget(budgetID uint, tokenID uuid.UUID) (*models.ListBudgetsResponse, error) {
	budget, err := bs.budgetRepository.GetBudgetById(budgetID)
	if err != nil {
		return nil, err
	}

	if budget.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return budget, nil
}

func (bs *BudgetService) GetBudgetJobs(budgetID uint, tokenID uuid.UUID) ([]models.Job, error) {
	budget, err := bs.budgetRepository.GetBudgetById(budgetID)
	if err != nil {
		return nil, err
	}

	if budget.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return bs.budgetRepository.GetBudgetJobs(budgetID)
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
