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

	if service.Duration < 0 {
		return utils.InvalidDurationError
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

func (ss *ServiceService) GetService(serviceID uint, tokenID uuid.UUID) (*models.Service, error) {
	service, err := ss.serviceRepository.GetServiceById(serviceID)
	if err != nil {
		return nil, err
	}

	if service.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return service, nil
}

func (ss *ServiceService) UpdateService(service *models.UpdateServiceRequest, tokenID uuid.UUID) (*models.Service, error) {
	if service.UserID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	if service.Name != nil && !utils.IsValidName(*service.Name) {
		return nil, utils.InvalidNameError
	}

	if service.Duration != nil && *service.Duration <= 0 {
		return nil, utils.InvalidDurationError
	}

	if service.Price != nil && *service.Price <= 0 {
		return nil, utils.InvalidPriceError
	}

	existingService, err := ss.serviceRepository.GetServiceById(service.ID)
	if err != nil {
		return nil, err
	}
	if existingService == nil {
		return nil, utils.NotFoundError
	}

	if service.Name != nil {
		capitalizedName := utils.CapitalizeName(*service.Name)
		service.Name = &capitalizedName
	}

	updatedService, err := ss.serviceRepository.Update(service)
	if err != nil {
		return nil, err
	}

	return updatedService, nil
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
