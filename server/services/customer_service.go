package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type CustomerService struct {
	customerRepository *repositories.CustomerRepository
}

func NewCustomerService(customerRepository *repositories.CustomerRepository) *CustomerService {
	return &CustomerService{customerRepository}
}

func (cs *CustomerService) CreateCustomer(customer *models.CreateCustomerRequest, tokenID uuid.UUID) error {
	if customer.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	formattedCPF, err := utils.FormatCPF(customer.CPF)
	if err != nil {
		return err
	}
	formattedName, err := utils.FormatName(customer.Name)
	if err != nil {
		return err
	}
	formattedEmail, err := utils.FormatEmail(customer.Email)
	if err != nil {
		return err
	}
	formattedPhone, err := utils.FormatPhone(customer.Phone)
	if err != nil {
		return err
	}

	validCustomer := &models.Customer{
		CPF:    formattedCPF,
		Name:   formattedName,
		Email:  formattedEmail,
		Phone:  formattedPhone,
		UserID: customer.UserID,
	}

	return cs.customerRepository.Create(validCustomer)
}

func (cs *CustomerService) ListCustomers(userID uuid.UUID) ([]models.Customer, error) {
	return cs.customerRepository.List(userID)
}

func (cs *CustomerService) GetCustomer(customerID uint, tokenID uuid.UUID) (*models.Customer, error) {
	customer, err := cs.customerRepository.GetCustomerById(customerID)
	if err != nil {
		return nil, err
	}

	if customer.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return customer, nil
}

func (cs *CustomerService) UpdateCustomer(customer *models.UpdateCustomerRequest, tokenID uuid.UUID) (*models.Customer, error) {	
	if customer.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	existingCustomer, err := cs.customerRepository.GetCustomerById(customer.ID)
	if err != nil {
		return nil, err
	}
	if existingCustomer == nil {
		return nil, utils.NotFoundError
	}

	if customer.CPF != nil {
		formattedCPF, err := utils.FormatCPF(*customer.CPF)
		if err != nil {
			return nil, err
		}
		customer.CPF = &formattedCPF
	}

	if customer.Name != nil {
		formattedName, err := utils.FormatName(*customer.Name)
		if err != nil {
			return nil, err
		}
		customer.Name = &formattedName
	}

	if customer.Email != nil {
		formattedEmail, err := utils.FormatEmail(*customer.Email)
		if err != nil {
			return nil, err
		}
		customer.Email = &formattedEmail
	}

	if customer.Phone != nil {
		formattedPhone, err := utils.FormatPhone(*customer.Phone)
		if err != nil {
			return nil, err
		}
		customer.Phone = &formattedPhone
	}

	updatedCustomer, err := cs.customerRepository.Update(customer)
	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (cs *CustomerService) DeleteCustomer(customerID uint, tokenID uuid.UUID) error {
	existingCustomer, err := cs.customerRepository.GetCustomerById(customerID)
	if err != nil {
		return err
	}

	if existingCustomer.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	return cs.customerRepository.Delete(customerID)
}
