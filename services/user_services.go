package services

import (
	"errors"
	"server/models"
	"server/repositories"
	"server/utils"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(user *models.User) (*models.User, error) {
	if !utils.IsValidEmail(user.Email) {
		return nil, errors.New("Invalid email")
	}

	if !utils.IsValidName(user.Name) {
		return nil, errors.New("Invalid name")
	}

	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("Email already in use")
	}

	return us.userRepository.Create(user)
}

func (us *UserService) ListUsers() ([]*models.User, error) {
	return us.userRepository.List()
}

func (us *UserService) GetUserById(id int) (*models.User, error) {
	user, err := us.userRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (us *UserService) UpdateUser(user *models.User) error {
	if !utils.IsValidEmail(user.Email) {
		return errors.New("Invalid email")
	}

	if !utils.IsValidName(user.Name) {
		return errors.New("Invalid name")
	}

	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return errors.New("Email already in use")
	}

	err = us.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(id int) error {
	err := us.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
