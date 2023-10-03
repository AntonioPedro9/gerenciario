package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type ClientService struct {
	clientRepository *repositories.ClientRepository
}

func NewClientService(clientRepository *repositories.ClientRepository) *ClientService {
	return &ClientService{clientRepository}
}

func (cs *ClientService) CreateClient(client *models.CreateClientRequest) error {
	if !utils.IsValidName(client.Name) {
		return utils.InvalidNameError
	}

	if !utils.IsValidEmail(client.Email) {
		return utils.InvalidEmailError
	}

	formattedNumber, err := utils.FormatPhone(client.Phone)
	if err != nil {
		return utils.InvalidPhoneError
	}

	validClient := &models.Client{
		Name:   utils.CapitalizeName(client.Name),
		Email:  client.Email,
		Phone:  formattedNumber,
		UserID: client.UserID,
	}

	return cs.clientRepository.Create(validClient)
}

func (cs *ClientService) ListClients(userID, tokenID uuid.UUID) ([]models.Client, error) {
	if userID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return cs.clientRepository.List(userID)
}

func (cs *ClientService) UpdateClient(client *models.UpdateClientRequest, tokenID uuid.UUID) error {
	if client.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	if !utils.IsValidName(client.Name) {
		return utils.InvalidNameError
	}

	if !utils.IsValidEmail(client.Email) {
		return utils.InvalidEmailError
	}

	existingClient, err := cs.clientRepository.GetClientById(client.ID)
	if err != nil {
		return err
	}
	if existingClient == nil {
		return utils.NotFoundError
	}

	formattedPhone, err := utils.FormatPhone(client.Phone)

	validClient := &models.UpdateClientRequest{
		ID:     client.ID,
		Name:   utils.CapitalizeName(client.Name),
		Email:  client.Email,
		Phone:  formattedPhone,
		UserID: client.UserID,
	}

	return cs.clientRepository.UpdateClient(validClient)
}

func (cs *ClientService) DeleteClient(clientID uint, authUserID uuid.UUID) error {
	existingClient, err := cs.clientRepository.GetClientById(clientID)
	if err != nil {
		return err
	}

	if existingClient.UserID != authUserID {
		return utils.UnauthorizedActionError
	}

	return cs.clientRepository.DeleteClient(clientID)
}