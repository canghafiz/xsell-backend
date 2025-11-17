package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(jwtKey string, duration time.Duration, data interface{}) (string, error) {
	if jwtKey == "" {
		return "", fmt.Errorf("JWT key is empty")
	}

	claims := jwt.MapClaims{
		"data": data,
		"exp":  time.Now().Add(duration).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtKey))
}
