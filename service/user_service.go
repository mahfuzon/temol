package service

import (
	"github.com/mahfuzon/temol/helper"
	"github.com/mahfuzon/temol/models"
	"github.com/mahfuzon/temol/repository"
	"github.com/mahfuzon/temol/request/user_request"
	"github.com/mahfuzon/temol/response/user_response"
	"github.com/sirupsen/logrus"
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
	log            *logrus.Logger
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository, log *logrus.Logger) *userService {
	return &userService{userRepository: userRepository, log: log}
}

func (userService *userService) Register(userRegisterRequest user_request.UserRegisterRequest) (user_response.UserProfileResponse, error) {
	userService.log.Info("userService.Register")
	user := models.User{
		Name:        userRegisterRequest.Name,
		Email:       helper.ConverseToSqlNullString(userRegisterRequest.Email),
		Password:    userRegisterRequest.Password,
		PhoneNumber: helper.ConverseToSqlNullString(userRegisterRequest.PhoneNumber),
		Address:     userRegisterRequest.Address,
	}

	user, err := userService.userRepository.Create(user)
	if err != nil {
		userService.log.Error(err.Error())
		return user_response.UserProfileResponse{}, err
	}

	userProfileResponse := user_response.ConverseToUserProfileResponse(user)

	userService.log.Info("success userService.Register")

	return userProfileResponse, nil
}

func (userService *userService) FindByEmailOrPhoneNumber(email, phoneNumber string) (models.User, error) {
	userService.log.Info("userService.FindByEmailOrPhoneNumber")
	user, err := userService.userRepository.FindByEmailOrPhoneNumber(email, phoneNumber)

	if err != nil {
		if err.Error() != "record not found" {
			return user, err
		}
		userService.log.Error(err.Error())
	}

	userService.log.Info("success userService.FindByEmailOrPhoneNumber")
	return user, nil
}

func (userService *userService) HashPassword(password string) (string, error) {
	userService.log.Info("userService.HashPassword")
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		userService.log.Error(err.Error())
		return string(passwordHash), err
	}

	userService.log.Info("success HashPassword")
	return string(passwordHash), nil
}

func (userService *userService) ValidatePassword(user models.User, password string) (bool, error) {
	userService.log.Info("ValidatePassword")
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		userService.log.Error(err.Error())
		return false, err
	}

	userService.log.Info("success ValidatePassword")
	return true, nil
}

func (userService *userService) FindById(userId int) (user_response.UserProfileResponse, error) {
	userService.log.Info("FindById")
	user, err := userService.userRepository.FindById(userId)
	if err != nil {
		userService.log.Error(err.Error())
		return user_response.UserProfileResponse{}, err
	}

	userResponse := user_response.ConverseToUserProfileResponse(user)
	userService.log.Info("success Find By Id")
	return userResponse, nil
}
