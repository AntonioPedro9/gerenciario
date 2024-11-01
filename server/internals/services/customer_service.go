package services

import (
	"server/internals/models"
	"server/internals/repositories"
)

type CustomerServiceInterface interface {
	CreateCustomer(data *models.CreateCustomerRequest, tokenUserId uint) error
	GetCustomer(customerId, tokenUserId uint) (models.GetCustomerResponse, error)
	GetUserCustomers(tokenUserId uint) (models.GetUserCustomersResponse, error)
	UpdateCustomer(customerId, tokenUserId uint, data *models.UpdateCustomerRequest) error
	DeleteCustomer(customerId, tokenUserId uint) error
}

type CustomerService struct {
	customerRepository repositories.CustomerRepositoryInterface
}

func NewCustomerService(customerRepository repositories.CustomerRepositoryInterface) *CustomerService {
	return &CustomerService{customerRepository}
}

func (cs *CustomerService) CreateCustomer(data *models.CreateCustomerRequest, tokenUserId uint) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return cs.customerRepository.Create(data, tokenUserId)
}

func (cs *CustomerService) GetCustomer(customerId, tokenUserId uint) (models.GetCustomerResponse, error) {
	var emptyCustomer models.GetCustomerResponse

	customer, err := cs.customerRepository.GetById(customerId, tokenUserId)
	if err != nil {
		return emptyCustomer, err
	}

	return *customer, nil
}

func (cs *CustomerService) GetUserCustomers(tokenUserId uint) (models.GetUserCustomersResponse, error) {
	var emptyCustomers models.GetUserCustomersResponse

	customers, err := cs.customerRepository.GetUserCustomers(tokenUserId)
	if err != nil {
		return emptyCustomers, err
	}

	customersArray := models.GetUserCustomersResponse{
		Customers: customers,
	}

	return customersArray, nil
}

func (cs *CustomerService) UpdateCustomer(customerId, tokenUserId uint, data *models.UpdateCustomerRequest) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return cs.customerRepository.Update(customerId, tokenUserId, data)
}

func (cs *CustomerService) DeleteCustomer(customerId, tokenUserId uint) error {
	return cs.customerRepository.Delete(customerId, tokenUserId)
}
