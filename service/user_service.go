package service

import (
	"github.com/mahfuzon/temol/helper"
	"github.com/mahfuzon/temol/models"
	"github.com/mahfuzon/temol/repository"
	"github.com/mahfuzon/temol/request/user_request"
	"github.com/mahfuzon/temol/response/user_response"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(userRegisterRequest user_request.UserRegisterRequest) (user_response.UserProfileResponse, error)
	FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error)
	HashPassword(password string) (string, error)
	ValidatePassword(user models.User, password string) (bool, error)
	FindById(userId int) (user_response.UserProfileResponse, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (userService *userService) Register(userRegisterRequest user_request.UserRegisterRequest) (user_response.UserProfileResponse, error) {
	user := models.User{
		Name:        userRegisterRequest.Name,
		Email:       helper.ConverseToSqlNullString(userRegisterRequest.Email),
		Password:    userRegisterRequest.Password,
		PhoneNumber: helper.ConverseToSqlNullString(userRegisterRequest.PhoneNumber),
		Address:     userRegisterRequest.Address,
	}

	user, err := userService.userRepository.Create(user)
	if err != nil {
		return user_response.UserProfileResponse{}, err
	}

	userProfileResponse := user_response.ConverseToUserProfileResponse(user)

	return userProfileResponse, nil
}

func (userService *userService) FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error) {
	user, err := userService.userRepository.FindByEmailOrPhoneNumber(email, phoneNumber)

	if err != nil {
		if err.Error() != "record not found" {
			return user, err
		}
	}

	return user, nil
}

func (userService *userService) HashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return string(passwordHash), err
	}

	return string(passwordHash), nil
}

func (userService *userService) ValidatePassword(user models.User, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (userService *userService) FindById(userId int) (user_response.UserProfileResponse, error) {
	user, err := userService.userRepository.FindById(userId)
	if err != nil {
		return user_response.UserProfileResponse{}, err
	}

	userResponse := user_response.ConverseToUserProfileResponse(user)
	return userResponse, nil
}
