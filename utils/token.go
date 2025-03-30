package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/om13rajpal/dbgpt/config"
)

func GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"iss": config.ISS,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	claimedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := claimedToken.SignedString([]byte(config.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return token, nil
}
