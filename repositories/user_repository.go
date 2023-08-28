package repositories

import (
	"server/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	result := ur.db.Create(user)

	if result.Error != nil {
		log.Error("Failed to create a new user: ", result.Error)
		return nil, result.Error
	}

	return user, nil
}

func (ur *UserRepository) List() ([]*models.User, error) {
	var users []*models.User

	result := ur.db.Find(&users)

	if result.Error != nil {
		log.Error("Failed to fetch users: ", result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (ur *UserRepository) GetUserById(id string) (*models.User, error) {
	var user models.User

	result := ur.db.First(&user, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Error("Failed to fetch user by ID: ", result.Error)

		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := ur.db.First(&user, "email = ?", email)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Error("Failed to fetch user by email: ", result.Error)

		return nil, result.Error
	}

	return &user, nil
}

func (ur *UserRepository) Update(user *models.UpdateUserRequest) error {
	result := ur.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name":     user.Name,
		"password": user.Password,
	})

	if result.Error != nil {
		log.Error("Failed to update user: ", result.Error)
		return result.Error
	}

	return nil
}

func (ur *UserRepository) Delete(id string) error {
	result := ur.db.Delete(&models.User{}, "id = ?", id)

	if result.Error != nil {
		log.Error("Failed to delete user: ", result.Error)
		return result.Error
	}

	return nil
}
