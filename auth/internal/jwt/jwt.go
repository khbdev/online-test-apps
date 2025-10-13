package jwt

import (
	"auth-service/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}


func GenerateTokens(username string) (*TokenPair, error) {
	now := time.Now()

	// --- ACCESS TOKEN ---
	accessClaims := jwt.MapClaims{
		"sub": username,
		"exp": now.Add(config.JWT.AccessTTL).Unix(),
		"iat": now.Unix(),
		"type": "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(config.JWT.AccessSecret))
	if err != nil {
		return nil, err
	}

	// --- REFRESH TOKEN ---
	refreshClaims := jwt.MapClaims{
		"sub": username,
		"exp": now.Add(config.JWT.RefreshTTL).Unix(),
		"iat": now.Unix(),
		"type": "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(config.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessStr,
		RefreshToken: refreshStr,
	}, nil
}
