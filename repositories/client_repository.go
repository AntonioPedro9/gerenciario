package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) *ClientRepository {
	return &ClientRepository{db}
}

func (cr *ClientRepository) Create(client *models.Client) error {
	return cr.db.Create(client).Error
}

func (cr *ClientRepository) List(tokenID uuid.UUID) ([]models.Client, error) {
	var clients []models.Client

	if err := cr.db.Where("user_id = ?", tokenID).Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

func (cr *ClientRepository) GetClientById(id uint) (*models.Client, error) {
	var client models.Client

	if err := cr.db.Where("id = ?", id).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &client, nil
}

func (cr *ClientRepository) UpdateClient(client *models.UpdateClientRequest) error {
	return cr.db.Model(&models.Client{}).
		Where("id = ?", client.ID).
		Updates(
			models.Client{
				Name:  client.Name,
				Email: client.Email,
				Phone: client.Phone,
			},
		).Error
}

func (cr *ClientRepository) DeleteClient(clientID uint) error {
	client := models.Client{ID: clientID}
	return cr.db.Delete(&client).Error
}
