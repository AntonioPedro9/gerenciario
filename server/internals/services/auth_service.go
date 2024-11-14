package services

import (
	"server/internals/models"
	"server/internals/repositories"
	"server/pkg/utils"

	"github.com/golang-jwt/jwt"
)

type AuthServiceInterface interface {
	Auth(data *models.LoginRequest) (string, string, error)
	Refresh(refreshTokenString string) (string, error)
}

type AuthService struct {
	userRepository repositories.UserRepositoryInterface
}

func NewAuthService(userRepository repositories.UserRepositoryInterface) *AuthService {
	return &AuthService{userRepository}
}

func (as *AuthService) Auth(data *models.LoginRequest) (string, string, error) {
	if err := validate.Struct(data); err != nil {
		return "", "", err
	}

	user, err := as.userRepository.GetByEmail(data.Email)
	if err != nil {
		return "", "", err
	}

	if err := utils.CheckPasswordHash(data.Password, user.Password); err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := utils.GenerateAccessAndRefreshToken(user.Id)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (as *AuthService) Refresh(refreshTokenString string) (string, error) {
	refreshToken, err := utils.VerifyRefreshToken(refreshTokenString)
	if err != nil {
		return "", err
	}

	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userId := uint(claims["sub"].(float64))

	newAccessToken, _, err := utils.GenerateAccessAndRefreshToken(userId)
	if err != nil {
		return "", err
	}

	return newAccessToken, nil
}
