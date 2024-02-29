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
	/**
	 * Data validations
	 **/
	if !utils.IsValidName(user.Name) {
		return utils.InvalidNameError
	}
	if !utils.IsValidEmail(user.Email) {
		return utils.InvalidEmailError
	}
	if len(user.Password) < 8 {
		return utils.WeakPasswordError
	}

	/**
	 * Checks if email is already in use
	 **/
	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return utils.EmailInUseError
	}

	/**
	 * Hashes password
	 **/
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	/**
	 * Generates UUID
	 **/
	userID, err := utils.GenerateUUID()
	if err != nil {
		return err
	}

	/**
	 * Create user
	 **/
	validUser := &models.User{
		ID:       userID,
		Name:     utils.CapitalizeText(user.Name),
		Email:    user.Email,
		Password: hashedPassword,
	}

	return us.userRepository.Create(validUser)
}

func (us *UserService) ListUsers() ([]models.User, error) {
	return us.userRepository.List()
}

func (us *UserService) UpdateUser(user *models.UpdateUserRequest, tokenID uuid.UUID) (*models.User, error) {
	/**
	 * Checks user authorization
	 **/
	if user.ID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	/**
	 * Data validations
	 **/
	if user.Name != nil && !utils.IsValidName(*user.Name) {
		return nil, utils.InvalidNameError
	}
	if user.Name != nil {
		capitalizedName := utils.CapitalizeText(*user.Name)
		user.Name = &capitalizedName
	}
	if user.Password != nil && len(*user.Password) < 8 {
		return nil, utils.WeakPasswordError
	}

	/**
	 * Checks if user exists
	 **/
	existingUser, err := us.userRepository.GetUserById(user.ID)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, utils.NotFoundError
	}

	/**
	 * Updates user
	 **/
	updatedUser, err := us.userRepository.Update(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (us *UserService) DeleteUser(id, tokenID uuid.UUID) error {
	/**
	 * Checks user authorization
	 **/
	if id != tokenID {
		return utils.UnauthorizedActionError
	}

	/**
	 * Checks if user exists
	 **/
	existingUser, err := us.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return utils.NotFoundError
	}

	/**
	 * Deletes user
	 **/
	return us.userRepository.Delete(id)
}

func (us *UserService) LoginUser(loginUserRequest *models.LoginUserResquest) (string, error) {
	/**
	 * Checks if user exists
	 **/
	existingUser, err := us.userRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil || existingUser == nil {
		return "", utils.InvalidEmailOrPasswordError
	}

	/**
	 * Compares password hashes
	 **/
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(loginUserRequest.Password)); err != nil {
		return "", utils.InvalidEmailOrPasswordError
	}

	/**
	 * Generates JWT token
	 **/
	tokenString, err := utils.GenerateToken(existingUser.ID)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

