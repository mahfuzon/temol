package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/mahfuzon/temol/libraries"
	"github.com/sirupsen/logrus"
	"os"
)

type AuthService interface {
	GenerateAccessToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
	log *logrus.Logger
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) GenerateAccessToken(userID int) (string, error) {
	s.log.Info("authService.GenerateAccessToken")
	token, err := libraries.GenerateNewToken(userID)

	if err != nil {
		s.log.Error(err.Error())
		return token, err
	}

	s.log.Info("success authService.GenerateAccessToken")
	return token, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	s.log.Info("authService.ValidateToken")
	secretKey := os.Getenv("jwt_secret_key")

	token, err := libraries.VerifyTokenBySecretKey(encodedToken, secretKey)

	if err != nil {
		s.log.Error(err.Error())
		return token, err
	}

	s.log.Info("success authService.ValidateToken")
	return token, nil
}
