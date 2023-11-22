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

func (cr *ClientRepository) List(userID uuid.UUID) ([]models.Client, error) {
	var clients []models.Client

	if err := cr.db.Where("user_id = ?", userID).Find(&clients).Error; err != nil {
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

func (cr *ClientRepository) Update(client *models.UpdateClientRequest) (*models.Client, error) {
	updateData := make(map[string]interface{})

	if client.CPF != nil {
		updateData["cpf"] = *client.CPF
	}

	if client.Name != nil {
		updateData["name"] = *client.Name
	}

	if client.Email != nil {
		updateData["email"] = *client.Email
	}

	if client.Phone != nil {
		updateData["phone"] = *client.Phone
	}

	err := cr.db.Model(&models.Client{}).
		Where("id = ?", client.ID).
		Updates(updateData).
		Error

	if err != nil {
		return nil, err
	}

	updatedClient := &models.Client{}
	err = cr.db.Where("id = ?", client.ID).First(updatedClient).Error
	if err != nil {
		return nil, err
	}

	return updatedClient, nil
}

func (cr *ClientRepository) Delete(clientID uint) error {
	client := models.Client{ID: clientID}
	return cr.db.Delete(&client).Error
}
