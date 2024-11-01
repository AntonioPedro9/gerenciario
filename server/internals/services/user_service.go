package services

import (
	"server/internals/models"
	"server/internals/repositories"
	"server/pkg/utils"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type UserServiceInterface interface {
	CreateUser(data *models.CreateUserRequest) error
	GetUser(tokenUserId uint) (models.GetUserResponse, error)
	UpdateUserData(data *models.UpdateUserDataRequest, tokenUserId uint) error
	UpdateUserPassword(data *models.UpdateUserPasswordRequest, tokenUserId uint) error
	DeleteUser(tokenUserId uint) error
}

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserService(userRepository repositories.UserRepositoryInterface) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(data *models.CreateUserRequest) error {
	if err := validate.Struct(data); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return err
	}
	data.Password = hashedPassword

	return us.userRepository.Create(data)
}

func (us *UserService) GetUser(tokenUserId uint) (models.GetUserResponse, error) {
	var emptyUser models.GetUserResponse

	user, err := us.userRepository.GetById(tokenUserId)
	if err != nil {
		return emptyUser, err
	}

	return *user, nil
}

func (us *UserService) UpdateUserData(data *models.UpdateUserDataRequest, tokenUserId uint) error {
	if err := validate.Struct(data); err != nil {
		return err
	}
	return us.userRepository.UpdateData(data, tokenUserId)
}

func (us *UserService) UpdateUserPassword(data *models.UpdateUserPasswordRequest, tokenUserId uint) error {
	if err := validate.Struct(data); err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		return err
	}

	validatedData := &models.UpdateUserPasswordRequest{
		Password: hashedPassword,
	}

	return us.userRepository.UpdatePassword(validatedData, tokenUserId)
}

func (us *UserService) DeleteUser(tokenUserId uint) error {
	return us.userRepository.Delete(tokenUserId)
}
