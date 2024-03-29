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
	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return utils.EmailInUseError
	}

	formattedName, err := utils.FormatName(user.Name)
	if err != nil {
		return err
	}
	formattedEmail, err := utils.FormatEmail(user.Email)
	if err != nil {
		return err
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
		Name:     formattedName,
		Email:    formattedEmail,
		Password: hashedPassword,
	}

	return us.userRepository.Create(validUser)
}

func (us *UserService) ListUsers() ([]models.ListUserResponse, error) {
	return us.userRepository.List()
}

func (us *UserService) GetUser(userID uuid.UUID, tokenID uuid.UUID) (*models.ListUserResponse, error) {
	user, err := us.userRepository.GetUserById(userID)
	if err != nil {
		return nil, err
	}

	if user.ID != tokenID {
		return nil, utils.UnauthorizedActionError
	}

	return user, nil
}

func (us *UserService) UpdateUser(user *models.UpdateUserRequest, tokenID uuid.UUID) error {
	if user.ID != tokenID {
		return utils.UnauthorizedActionError
	}

	existingUser, err := us.userRepository.GetUserById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return utils.NotFoundError
	}

	if user.Name != nil {
		formattedName, err := utils.FormatName(*user.Name)
		if err != nil {
			return err
		}
		user.Name = &formattedName
	}
	
	if user.Password != nil {
		hashedPassword, err := utils.HashPassword(*user.Password)
		if err != nil {
			return err
		}
		user.Password = &hashedPassword
	}

	if err := us.userRepository.Update(user); err != nil {
		return err
	}

	return nil
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

	return us.userRepository.Delete(id)
}

func (us *UserService) LoginUser(loginUserRequest *models.LoginUserRequest) (string, error) {
	existingUser, err := us.userRepository.GetUserByEmail(loginUserRequest.Email)
	if err != nil || existingUser == nil {
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
