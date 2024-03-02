package repositories

import (
	"server/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) List() ([]models.ListUserResponse, error) {
	var users []models.ListUserResponse

	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserById(id uuid.UUID) (*models.ListUserResponse, error) {
	var user models.ListUserResponse

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	var count int64

	ur.db.Model(&user).Where("email = ?", email).Count(&count)
	if count == 0 {
		return nil, nil
	}

	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Update(user *models.UpdateUserRequest) error {
	err := ur.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) Delete(id uuid.UUID) error {
	user := models.User{ID: id}
	return ur.db.Delete(&user).Error
}
