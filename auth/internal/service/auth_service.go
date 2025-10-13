package service

import (
	"auth-service/internal/client"
	"auth-service/internal/jwt"
	"errors"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Login(username, password string) (*jwt.TokenPair, error) {
	if !client.FindUser(username, password) {
		return nil, errors.New("invalid username or password")
	}

	return jwt.GenerateTokens(username)
}
