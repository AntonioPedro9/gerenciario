package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{db}
}

func (sr *ServiceRepository) Create(service *models.Service) error {
	return sr.db.Create(service).Error
}

func (sr *ServiceRepository) List(userID uuid.UUID) ([]models.Service, error) {
	var services []models.Service

	if err := sr.db.Where("user_id = ?", userID).Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

func (sr *ServiceRepository) GetServiceById(id uint) (*models.Service, error) {
	var service models.Service

	if err := sr.db.Where("id = ?", id).First(&service).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &service, nil
}

func (sr *ServiceRepository) Update(service *models.UpdateServiceRequest) (*models.Service, error) {
	updateData := make(map[string]interface{})

	if service.Name != nil {
		updateData["name"] = *service.Name
	}

	if service.Description != nil {
		updateData["description"] = *service.Description
	}

	if service.Duration != nil {
		updateData["duration"] = *service.Duration
	}

	if service.Price != nil {
		updateData["price"] = *service.Price
	}

	err := sr.db.Model(&models.Service{}).
		Where("id = ?", service.ID).
		Updates(updateData).
		Error

	if err != nil {
		return nil, err
	}

	updatedService := &models.Service{}
	err = sr.db.Where("id = ?", service.ID).First(updatedService).Error
	if err != nil {
		return nil, err
	}

	return updatedService, nil
}

func (sr *ServiceRepository) Delete(serviceID uint) error {
	service := models.Service{ID: serviceID}
	return sr.db.Delete(&service).Error
}
