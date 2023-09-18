package services

import (
	"errors"
	"server/models"
	"server/repositories"
	"server/utils"

	"github.com/google/uuid"
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
		Name:     utils.CapitalizeName(user.Name),
		Email:    user.Email,
		Password: utils.HashPassword(user.Password),
	}

	return us.userRepository.Create(validUser)
}

func (us *UserService) ListUsers() ([]models.User, error) {
	return us.userRepository.List()
}

func (us *UserService) UpdateUser(user *models.UpdateUserRequest) error {
	// pegar o token pelo cookie

	// parsear o token

	// extrair o id de usuário do token

	// verificar se o id do token é igual o id do usuário que fez a requisição

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

func (us *UserService) DeleteUser(id uuid.UUID) error {
	// pegar o token pelo cookie

	// parsear o token

	// extrair o id de usuário do token

	// verificar se o id do token é igual o id do usuário que fez a requisição

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

	tokenString, err := utils.GenerateToken(existingUser.ID.String())
	if err != nil {
		return "", errors.New("Failed to generate JWT token")
	}

	return tokenString, nil
}
