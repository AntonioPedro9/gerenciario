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

func (sr *ServiceRepository) List(tokenID uuid.UUID) ([]models.Service, error) {
	var services []models.Service

	if err := sr.db.Where("user_id = ?", tokenID).Find(&services).Error; err != nil {
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

func (sr *ServiceRepository) Update(service *models.Service) error {
	return sr.db.Model(&models.Service{}).
		Where("id = ?", service.ID).
		Updates(
			models.Service{
				Name:        service.Name,
				Description: service.Description,
				Duration:    service.Duration,
				Price:       service.Price,
			},
		).Error
}

func (sr *ServiceRepository) Delete(serviceID uint) error {
	service := models.Service{ID: serviceID}
	return sr.db.Delete(&service).Error
}
