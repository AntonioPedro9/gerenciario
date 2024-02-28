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

func (cs *CustomerService) CreateCustomer(customer *models.CreateCustomerRequest) error {
	var formattedCPF string
	var err error
	
	if customer.CPF != "" {
		formattedCPF, err = utils.FormatCPF(customer.CPF)
		if err != nil {
			return err
		}
	}

	if !utils.IsValidName(customer.Name) {
		return utils.InvalidNameError
	}

	if customer.Email != "" {
		if !utils.IsValidEmail(customer.Email) {
			return utils.InvalidEmailError
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

func (cs *CustomerService) ListCustomers(userID, tokenID uuid.UUID) ([]models.Customer, error) {
	if userID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

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
			return nil, utils.InvalidCpfError
		}
		customer.CPF = &formattedCPF
	}

	if customer.Name != nil {
		if !utils.IsValidName(*customer.Name) {
			return nil, utils.InvalidNameError
		}
		capitalizedName := utils.CapitalizeText(*customer.Name)
		customer.Name = &capitalizedName
	}

	if customer.Email != nil {
		if !utils.IsValidEmail(*customer.Email) {
			return nil, utils.InvalidEmailError
		}
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
