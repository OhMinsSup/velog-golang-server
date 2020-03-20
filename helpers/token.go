package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var signingKey = []byte("AllYourBase")

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
		return "", ErrorGenerateToken
	}

	return tokenString, nil
}

func DecodeToken(deocedToken string) (JSON, error) {
	result, err := jwt.Parse(deocedToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method: %v", token.Header["alg"])
			return nil, ErrorSigningMethod
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, ErrorInvalidToken
	}

	if !result.Valid {
		return nil, ErrorInvalidToken
	}
	return result.Claims.(jwt.MapClaims), nil
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
