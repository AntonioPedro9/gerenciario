package services

import (
	"server/models"
	"server/repositories"
	"server/utils"
	"server/utils/validations"

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
	
	err := validations.ValidateCreateCustomerRequest(customer)
	if err != nil {
		return nil
	}

	var formattedCPF string
	if customer.CPF != "" {
		formattedCPF, err = utils.FormatCPF(customer.CPF)
		if err != nil {
			return err
		}
	}

	formattedPhone, err := utils.FormatPhone(customer.Phone)
	if err != nil {
		return utils.InvalidPhoneError
	}

	validCustomer := &models.Customer{
		CPF:    formattedCPF,
		Name:   utils.CapitalizeText(customer.Name),
		Email:  customer.Email,
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

	err = validations.ValidateUpdateCustomerRequest(customer)
	if err != nil {
		return nil, err
	}

	if customer.CPF != nil {
		formattedCPF, err := utils.FormatCPF(*customer.CPF)
		if err != nil {
			return nil, utils.InvalidCpfError
		}
		customer.CPF = &formattedCPF
	}

	if customer.Phone != nil {
		formattedPhone, err := utils.FormatPhone(*customer.Phone)
		if err != nil {
			return nil, utils.InvalidPhoneError
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
