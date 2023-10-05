package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OfferingRepository struct {
	db *gorm.DB
}

func NewOfferingRepository(db *gorm.DB) *OfferingRepository {
	return &OfferingRepository{db}
}

func (or *OfferingRepository) Create(offering *models.Offering) error {
	return or.db.Create(offering).Error
}

func (or *OfferingRepository) List(tokenID uuid.UUID) ([]models.Offering, error) {
	var offerings []models.Offering

	if err := or.db.Where("user_id = ?", tokenID).Find(&offerings).Error; err != nil {
		return nil, err
	}

	return offerings, nil
}

func (or *OfferingRepository) GetOfferingById(id uint) (*models.Offering, error) {
	var offering models.Offering

	if err := or.db.Where("id = ?", id).First(&offering).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &offering, nil
}

func (or *OfferingRepository) Update(offering *models.Offering) error {
	return or.db.Model(&models.Offering{}).
		Where("id = ?", offering.ID).
		Updates(
			models.Offering{
				Name:        offering.Name,
				Description: offering.Description,
				Duration:    offering.Duration,
			},
		).Error
}

func (or *OfferingRepository) Delete(offeringID uint) error {
	offering := models.Offering{ID: offeringID}
	return or.db.Delete(&offering).Error
}
