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
		return utils.InvalidNameError
	}

	if !utils.IsValidEmail(user.Email) {
		return utils.InvalidEmailError
	}

	if len(user.Password) < 8 {
		return utils.WeakPasswordError
	}

	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return utils.EmailInUseError
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	userID, err := utils.GenerateUUID()
	if err != nil {
		return err
	}

	validUser := &models.User{
		ID:       userID,
		Name:     utils.CapitalizeName(user.Name),
		Email:    user.Email,
		Password: hashedPassword,
	}

	return us.userRepository.Create(validUser)
}

func (us *UserService) ListUsers() ([]models.User, error) {
	return us.userRepository.List()
}

func (us *UserService) UpdateUser(user *models.UpdateUserRequest, tokenID uuid.UUID) error {
	if user.ID != tokenID {
		return utils.UnauthorizedActionError
	}

	if !utils.IsValidName(user.Name) {
		return utils.InvalidNameError
	}

	existingUser, err := us.userRepository.GetUserById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return utils.NotFoundError
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	validUser := &models.UpdateUserRequest{
		ID:       user.ID,
		Name:     utils.CapitalizeName(user.Name),
		Password: hashedPassword,
	}

	return us.userRepository.UpdateUser(validUser)
}

func (us *UserService) DeleteUser(id, tokenID uuid.UUID) error {
	if id != tokenID {
		return utils.UnauthorizedActionError
	}

	existingUser, err := us.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return utils.NotFoundError
	}

	return us.userRepository.DeleteUser(id)
}

func (us *UserService) LoginUser(loginUserRequest *models.LoginUserResquest) (string, error) {
	existingUser, err := us.userRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil {
		return "", utils.InvalidEmailOrPasswordError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginUserRequest.Password)); err != nil {
		return "", utils.InvalidEmailOrPasswordError
	}

	tokenString, err := utils.GenerateToken(existingUser.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
