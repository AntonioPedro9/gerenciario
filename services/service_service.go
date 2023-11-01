package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type ServiceService struct {
	serviceRepository *repositories.ServiceRepository
}

func NewServiceService(serviceRepository *repositories.ServiceRepository) *ServiceService {
	return &ServiceService{serviceRepository}
}

func (ss *ServiceService) CreateService(service *models.CreateServiceRequest) error {
	if !utils.IsValidName(service.Name) {
		return utils.InvalidNameError
	}

	if service.Price < 0 {
		return utils.InvalidPriceError
	}

	validService := &models.Service{
		Name:        utils.CapitalizeName(service.Name),
		Description: service.Description,
		Duration:    service.Duration,
		Price:       service.Price,
		UserID:      service.UserID,
	}

	return ss.serviceRepository.Create(validService)
}

func (ss *ServiceService) ListServices(userID, tokenID uuid.UUID) ([]models.Service, error) {
	if userID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return ss.serviceRepository.List(userID)
}

func (ss *ServiceService) UpdateService(service *models.UpdateServiceRequest, tokenID uuid.UUID) error {
	if service.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	if !utils.IsValidName(service.Name) {
		return utils.InvalidNameError
	}

	existingService, err := ss.serviceRepository.GetServiceById(service.ID)
	if err != nil {
		return err
	}
	if existingService == nil {
		return utils.NotFoundError
	}

	validService := &models.Service{
		ID:          service.ID,
		Name:        utils.CapitalizeName(service.Name),
		Description: service.Description,
		Duration:    service.Duration,
		Price:       service.Price,
		UserID:      service.UserID,
	}

	return ss.serviceRepository.Update(validService)
}

func (ss *ServiceService) DeleteService(serviceID uint, tokenID uuid.UUID) error {
	existingService, err := ss.serviceRepository.GetServiceById(serviceID)
	if err != nil {
		return err
	}

	if existingService.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	return ss.serviceRepository.Delete(serviceID)
}
