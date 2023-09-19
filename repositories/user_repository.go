package repositories

import (
	"server/models"

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

func (ur *UserRepository) List() ([]models.User, error) {
	var users []models.User

	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *models.UpdateUserRequest) error {
	return ur.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{Name: user.Name, Password: user.Password}).Error
}

func (ur *UserRepository) DeleteUser(id string) error {
	user := models.User{ID: id}
	return ur.db.Delete(&user).Error
}
