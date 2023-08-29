package services

import (
	"errors"
	"os"
	"server/models"
	"server/repositories"
	"server/utils"
	"time"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
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
		return nil, errors.New("Name can only contain letters and must be at least 2 characters long")
	}

	user.Name = utils.CapitalizeName(user.Name)

	existingUser, err := us.userRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("User already exists")
	}

	log.Info("Creating user")
	return us.userRepository.Create(user)
}

func (us *UserService) ListUsers() ([]*models.User, error) {
	log.Info("Listing users")
	return us.userRepository.List()
}

func (us *UserService) GetUserById(id string) (*models.User, error) {
	user, err := us.userRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}

	log.Info("Getting user by id")
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

	log.Info("Getting user by email")
	return user, nil
}

func (us *UserService) UpdateUser(user *models.UpdateUserRequest) error {
	if !utils.IsValidName(user.Name) {
		return errors.New("Invalid name")
	}

	user.Name = utils.CapitalizeName(user.Name)

	existingUser, err := us.userRepository.GetUserById(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("User not found")
	}

	err = us.userRepository.Update(user)
	if err != nil {
		return err
	}

	log.Info("Updating user")
	return nil
}

func (us *UserService) DeleteUser(id string) error {
	user, err := us.userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User not found")
	}

	err = us.userRepository.Delete(id)
	if err != nil {
		return err
	}

	log.Info("Deleting user")
	return nil
}

func (us *UserService) Login(email, password string) (string, error) {
	user, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	log.Info("Logging in")
	return tokenString, nil
}
