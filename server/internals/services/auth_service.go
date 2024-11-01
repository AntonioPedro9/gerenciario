package services

import (
	"server/internals/models"
	"server/internals/repositories"
	"server/pkg/utils"
)

type AuthServiceInterface interface {
	Auth(data *models.LoginRequest) (string, error)
}

type AuthService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewAuthService(userRepository repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{userRepository}
}

func (as *AuthService) Auth(data *models.LoginRequest) (string, error) {
	if err := validate.Struct(data); err != nil {
		return "", err
	}

	user, err := as.userRepository.GetByEmail(data.Email)
	if err != nil {
		return "", err
	}

	if err := utils.CheckPasswordHash(data.Password, user.Password); err != nil {
		return "", err
	}

	tokenString, err := utils.GenerateToken(user.Id)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
