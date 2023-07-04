package services

import (
	"errors"
	"server/models"
	"server/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(user *models.User) error {
	existingUser, err := us.userRepository.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("Email already in use")
	}

	return us.userRepository.Create(user)
}

func (us *UserService) ListUsers() ([]*models.User, error) {
	return us.userRepository.List()
}

func (us *UserService) UpdateUser(user *models.User) error {
	existingUser, err := us.userRepository.FindByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return errors.New("Email already in use")
	}

	return us.userRepository.Update(user)
}

func (us *UserService) DeleteUser(id int) error {
	return us.userRepository.Delete(id)
}
