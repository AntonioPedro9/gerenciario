package repositories

import (
	"server/internals/models"
	"server/pkg/errors"
	"server/pkg/logs"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(data *models.CreateUserRequest) error
	GetById(id uint) (*models.GetUserResponse, error)
	GetByEmail(email string) (*models.User, error)
	UpdateData(data *models.UpdateUserDataRequest, id uint) error
	UpdatePassword(data *models.UpdateUserPasswordRequest, id uint) error
	Delete(id uint) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(data *models.CreateUserRequest) error {
	user := &models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	result := ur.db.Create(user)

	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return errors.EmailInUseError
		}
		logs.LogError(result.Error)
		return result.Error
	}

	return nil
}

func (ur *UserRepository) GetById(id uint) (*models.GetUserResponse, error) {
	var user models.User

	result := ur.db.First(&user, id)

	if result.RowsAffected == 0 {
		return nil, errors.NotFoundError
	}

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	return &models.GetUserResponse{
		Id:    uint(user.Id),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	result := ur.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		logs.LogError(result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.NotFoundError
	}

	return &models.User{
		Id:       uint(user.Id),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (ur *UserRepository) UpdateData(data *models.UpdateUserDataRequest, id uint) error {
	user := &models.User{
		Id:   id,
		Name: data.Name,
	}

	result := ur.db.Updates(user)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}

func (ur *UserRepository) UpdatePassword(data *models.UpdateUserPasswordRequest, id uint) error {
	result := ur.db.Model(&models.User{}).Where("id = ?", id).Update("password", data.Password)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}

func (ur *UserRepository) Delete(id uint) error {
	result := ur.db.Delete(&models.User{}, id)

	if result.Error != nil {
		logs.LogError(result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.NotFoundError
	}

	return nil
}
