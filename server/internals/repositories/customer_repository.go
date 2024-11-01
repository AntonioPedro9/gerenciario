package repositories

import (
	"server/internals/models"
	"server/pkg/errors"
	"server/pkg/logs"

	"gorm.io/gorm"
)

type CustomerRepositoryInterface interface {
	Create(data *models.CreateCustomerRequest, tokenUserId uint) error
	GetById(customerId, tokenUserId uint) (*models.GetCustomerResponse, error)
	GetUserCustomers(tokenUserId uint) ([]models.GetCustomerResponse, error)
	Update(customerId, tokenUserId uint, data *models.UpdateCustomerRequest) error
	Delete(customerId, tokenUserId uint) error
}

type CustomerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{db}
}

func (cr *CustomerRepository) Create(data *models.CreateCustomerRequest, tokenUserId uint) error {
	customer := &models.Customer{
		UserId: tokenUserId,
		Name:   data.Name,
		CPF:    data.CPF,
		Phone:  data.Phone,
		Email:  data.Email,
	}

	result := cr.db.Create(customer)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	return nil
}

func (cr *CustomerRepository) GetById(customerId, tokenUserId uint) (*models.GetCustomerResponse, error) {
	var customer models.Customer

	result := cr.db.Where("id = ? AND user_id = ?", customerId, tokenUserId).First(&customer)

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFoundError
	}

	return &models.GetCustomerResponse{
		Id:     customer.Id,
		UserId: customer.UserId,
		Name:   customer.Name,
		CPF:    customer.CPF,
		Phone:  customer.Phone,
		Email:  customer.Email,
	}, nil
}

func (cr *CustomerRepository) GetUserCustomers(tokenUserId uint) ([]models.GetCustomerResponse, error) {
	var customers []models.Customer
	var customerResponses []models.GetCustomerResponse

	result := cr.db.Where("user_id = ?", tokenUserId).Find(&customers)

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	for _, customer := range customers {
		customerResponses = append(customerResponses, models.GetCustomerResponse{
			Id:     customer.Id,
			UserId: customer.UserId,
			Name:   customer.Name,
			CPF:    customer.CPF,
			Phone:  customer.Phone,
			Email:  customer.Email,
		})
	}

	return customerResponses, nil
}

func (cr *CustomerRepository) Update(customerId, tokenUserId uint, data *models.UpdateCustomerRequest) error {
	customer := &models.Customer{
		Name:  data.Name,
		Phone: data.Phone,
		CPF:   data.CPF,
		Email: data.Email,
	}

	result := cr.db.Model(&models.Customer{}).Where("id = ? AND user_id = ?", customerId, tokenUserId).Updates(customer)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}

func (cr *CustomerRepository) Delete(customerId, tokenUserId uint) error {
	result := cr.db.Where("id = ? AND user_id = ?", customerId, tokenUserId).Delete(&models.Customer{})

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}
