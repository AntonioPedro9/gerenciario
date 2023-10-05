package services

import (
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
)

type OfferingService struct {
	offeringRepository *repositories.OfferingRepository
}

func NewOfferingService(offeringRepository *repositories.OfferingRepository) *OfferingService {
	return &OfferingService{offeringRepository}
}

func (os *OfferingService) CreateOffering(offering *models.CreateOfferingRequest) error {
	if !utils.IsValidName(offering.Name) {
		return utils.InvalidNameError
	}

	validOffering := &models.Offering{
		Name:        utils.CapitalizeName(offering.Name),
		Description: offering.Description,
		Duration:    offering.Duration,
		UserID:      offering.UserID,
	}

	return os.offeringRepository.Create(validOffering)
}

func (os *OfferingService) ListOfferings(userID, tokenID uuid.UUID) ([]models.Offering, error) {
	if userID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return os.offeringRepository.List(userID)
}

func (os *OfferingService) UpdateOffering(offering *models.UpdateOfferingRequest, tokenID uuid.UUID) error {
	if offering.UserID != tokenID {
		return utils.UnauthorizedActionError
	}

	if !utils.IsValidName(offering.Name) {
		return utils.InvalidNameError
	}

	existingOffering, err := os.offeringRepository.GetOfferingById(offering.ID)
	if err != nil {
		return err
	}
	if existingOffering == nil {
		return utils.NotFoundError
	}

	validOffering := &models.Offering{
		ID:          offering.ID,
		Name:        utils.CapitalizeName(offering.Name),
		Description: offering.Description,
		Duration:    offering.Duration,
		UserID:      offering.UserID,
	}

	return os.offeringRepository.Update(validOffering)
}

func (os *OfferingService) DeleteOffering(offeringID uint, authUserID uuid.UUID) error {
	existingOffering, err := os.offeringRepository.GetOfferingById(offeringID)
	if err != nil {
		return err
	}

	if existingOffering.UserID != authUserID {
		return utils.UnauthorizedActionError
	}

	return os.offeringRepository.Delete(offeringID)
}
