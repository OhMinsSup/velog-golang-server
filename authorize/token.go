package authorize

import (
	"errors"
	"github.com/OhMinsSup/story-server/ent"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"log"
	"time"
)

var (
	signingKey       = []byte("AllYourBase")
	accessTokenName  = "access_token"
	refreshTokenName = "refresh_token"
)

func GenerateUserToken(user *ent.User, authToken *ent.AuthToken) (string, string) {
	accessSubject := "access_token"
	accessPayload := libs.JSON{
		"user_id": user.ID,
	}

	refreshSubject := "refresh_token"
	refreshPayload := libs.JSON{
		"user_id":  user.ID,
		"token_id": authToken.ID,
	}

	accessToken, _ := GenerateAccessToken(accessPayload, accessSubject)
	refreshToken, _ := GenerateRefreshToken(refreshPayload, refreshSubject)

	return accessToken, refreshToken
}

func RefreshUserToken(user *ent.User, tokenId uuid.UUID, originalRefreshToken string, refreshTokenExp int64) (string, string) {
	now := time.Now().Unix()
	diff := refreshTokenExp - now

	refreshToken := originalRefreshToken
	accessSubject := "access_token"
	accessPayload := libs.JSON{
		"user_id": user.ID,
	}
	accessToken, _ := GenerateAccessToken(accessPayload, accessSubject)
	if diff < 60*60*24*15 {
		log.Println("refreshing....")
		refreshSubject := "refresh_token"
		refreshPayload := libs.JSON{
			"user_id":  user.ID,
			"token_id": tokenId,
		}

		refreshToken, _ = GenerateRefreshToken(refreshPayload, refreshSubject)
	}

	return accessToken, refreshToken
}

func generateToken(payload libs.JSON, subject string, expire time.Duration) (string, error) {
	// Create the Claims
	claims := &jwt.MapClaims{
		"exp":     time.Now().Add(expire).Unix(),
		"issuer":  "storeis.vercel.app",
		"subject": subject,
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", errors.New("Token Signed Error")
	}

	return tokenString, nil
}

func DecodeToken(deocedToken string) (libs.JSON, error) {
	result, err := jwt.Parse(deocedToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("Unexpected signing method: %v", token.Header["alg"])
			return nil, errors.New("Token Signed Error")
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, errors.New("Token Invalid Error")
	}

	if !result.Valid {
		return nil, errors.New("Token Invalid Error")
	}

	return result.Claims.(jwt.MapClaims), nil
}

func GenerateRegisterToken(payload libs.JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24)
}

func GenerateAccessToken(payload libs.JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24)
}

func GenerateRefreshToken(payload libs.JSON, subject string) (string, error) {
	return generateToken(payload, subject, time.Hour*24*30)
}
