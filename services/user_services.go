package services

import (
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
		return models.InvalidNameError
	}

	if !utils.IsValidEmail(user.Email) {
		return models.InvalidEmailError
	}

	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return models.EmailInUseError
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

func (us *UserService) UpdateUser(user *models.UpdateUserRequest, tokenID uuid.UUID) error {
	if user.ID != tokenID {
		return models.UnauthorizedActionError
	}

	if !utils.IsValidName(user.Name) {
		return models.InvalidNameError
	}

	existingUser, err := us.userRepository.GetUserById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return models.NotFoundError
	}

	validUser := &models.UpdateUserRequest{
		ID:       user.ID,
		Name:     utils.CapitalizeName(user.Name),
		Password: utils.HashPassword(user.Password),
	}

	return us.userRepository.UpdateUser(validUser)
}

func (us *UserService) DeleteUser(id, tokenID uuid.UUID) error {
	if id != tokenID {
		return models.UnauthorizedActionError
	}

	existingUser, err := us.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return models.EmailInUseError
	}

	return us.userRepository.DeleteUser(id)
}

func (us *UserService) LoginUser(loginUserRequest *models.LoginUserResquest) (string, error) {
	existingUser, err := us.userRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil {
		return "", models.InvalidEmailOrPasswordError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginUserRequest.Password)); err != nil {
		return "", models.InvalidEmailOrPasswordError
	}

	tokenString := utils.GenerateToken(existingUser.ID)

	return tokenString, nil
}
