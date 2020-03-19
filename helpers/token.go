package helpers

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var signingKey = []byte("AllYourBase")
var errorGenerateToken = errors.New("Generate Token Error")

func generateToken(payload JSON, subject string, expire time.Duration) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"exp":     time.Now().Add(expire).Unix(),
		"issuer":  "velog.io",
		"subject": subject,
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", errorGenerateToken
	}

	return tokenString, nil
}

func GenerateRegisterToken(payload JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24)
}

func GenerateAccessToken(payload JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24)
}

func GenerateRefreshToken(payload JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24*30)
}
