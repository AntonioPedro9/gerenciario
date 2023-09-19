package services

import (
	"errors"
	"server/models"
	"server/repositories"
	"server/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(user *models.CreateUserRequest) error {
	if !utils.IsValidName(user.Name) {
		return errors.New("Invalid name")
	}

	if !utils.IsValidEmail(user.Email) {
		return errors.New("Invalid email")
	}

	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("Email already in use")
	}

	validUser := &models.User{
		ID:       utils.GenerateUUID(),
		Name:     utils.CapitalizeName(user.Name),
		Email:    user.Email,
		Password: utils.HashPassword(user.Password),
	}

	return us.userRepository.Create(validUser)
}

func (us *UserService) ListUsers() ([]models.User, error) {
	return us.userRepository.List()
}

func (us *UserService) UpdateUser(user *models.UpdateUserRequest, tokenID string) error {
	if user.ID != tokenID {
		return errors.New("Unauthorized action")
	}

	if !utils.IsValidName(user.Name) {
		return errors.New("Invalid name")
	}

	existingUser, err := us.userRepository.GetUserById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("User not found")
	}

	validUser := &models.UpdateUserRequest{
		ID:       user.ID,
		Name:     utils.CapitalizeName(user.Name),
		Password: utils.HashPassword(user.Password),
	}

	return us.userRepository.UpdateUser(validUser)
}

func (us *UserService) DeleteUser(id, tokenID string) error {
	if id != tokenID {
		return errors.New("Unauthorized action")
	}

	existingUser, err := us.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("User not found")
	}

	return us.userRepository.DeleteUser(id)
}

func (us *UserService) LoginUser(loginUserRequest *models.LoginUserResquest) (string, error) {
	existingUser, err := us.userRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginUserRequest.Password)); err != nil {
		return "", errors.New("Invalid email or password")
	}

	tokenString := utils.GenerateToken(existingUser.ID)

	return tokenString, nil
}
