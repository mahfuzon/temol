package service

import (
	"github.com/golang-jwt/jwt"
	"github.com/mahfuzon/temol/libraries"
	"os"
)

type AuthService interface {
	GenerateAccessToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type authService struct {
}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) GenerateAccessToken(userID int) (string, error) {
	token, err := libraries.GenerateNewToken(userID)

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	secretKey := os.Getenv("jwt_secret_key")

	token, err := libraries.VerifyTokenBySecretKey(encodedToken, secretKey)

	if err != nil {
		return token, err
	}

	return token, nil
}
