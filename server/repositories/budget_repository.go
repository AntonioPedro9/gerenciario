package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) *BudgetRepository {
	return &BudgetRepository{db}
}

func (br *BudgetRepository) Create(budget *models.Budget) error {
	return br.db.Create(budget).Error
}

func (br *BudgetRepository) CreateBudgetJob(budgetJob *models.BudgetJob) error {
	return br.db.Create(budgetJob).Error
}

func (br *BudgetRepository) List(userID uuid.UUID) ([]models.ListBudgetsResponse, error) {
	var budgets []models.Budget
	var response []models.ListBudgetsResponse

	if err := br.db.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		return nil, err
	}

	for _, budget := range budgets {
		var customer models.Customer
		if err := br.db.Where("id = ?", budget.CustomerID).First(&customer).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			} else {
				return nil, err
			}
		}

		response = append(response, models.ListBudgetsResponse{
			UserID:        budget.UserID,
			ID:            budget.ID,
			CustomerID:    customer.ID,
			CustomerName:  customer.Name,
			CustomerPhone: customer.Phone,
			BudgetJobs:    budget.BudgetJobs,
			BudgetDate:    budget.BudgetDate,
			ScheduledDate: budget.ScheduledDate,
			Vehicle:       budget.Vehicle,
			LicensePlate:  budget.LicensePlate,
			Price:         budget.Price,
		})
	}

	return response, nil
}

func (br *BudgetRepository) GetBudgetById(id uint) (*models.ListBudgetsResponse, error) {
	var budget models.Budget
	var response models.ListBudgetsResponse

	if err := br.db.Where("id = ?", id).First(&budget).Error; err != nil {
		return nil, err
	}

	var customer models.Customer
	if err := br.db.Where("id = ?", budget.CustomerID).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	response = models.ListBudgetsResponse{
		UserID:        budget.UserID,
		ID:            budget.ID,
		CustomerID:    customer.ID,
		CustomerName:  customer.Name,
		CustomerPhone: customer.Phone,
		BudgetJobs:    budget.BudgetJobs,
		BudgetDate:    budget.BudgetDate,
		ScheduledDate: budget.ScheduledDate,
		Vehicle:       budget.Vehicle,
		LicensePlate:  budget.LicensePlate,
		Price:         budget.Price,
	}

	return &response, nil
}


func (br *BudgetRepository) GetBudgetJobs(budgetID uint) ([]models.Job, error) {
	var budgetJobs []models.BudgetJob
	var jobs []models.Job

	if err := br.db.Where("budget_id = ?", budgetID).Find(&budgetJobs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	for _, budgetJob := range budgetJobs {
		var job models.Job
		
		if err := br.db.Where("id = ?", budgetJob.JobID).First(&job).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return nil, err
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

func (br *BudgetRepository) Delete(budgetID uint) error {
	budget := models.Budget{ID: budgetID}
	return br.db.Delete(&budget).Error
}
