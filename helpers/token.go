package helpers

import (
	"context"
	"github.com/OhMinsSup/story-server/ent"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var (
	signingKey       = []byte("AllYourBase")
	accessTokenName  = "access_token"
	refreshTokenName = "refresh_token"
)

func GenerateUserToken(user *ent.User, client *ent.Client, bg context.Context) (string, string, error) {
	log.Println(user)

	tx, err := client.Tx(bg)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	authToken, err := tx.AuthToken.
		Create().
		Save(bg)

	if err != nil {
		log.Println("auth token generate error ::", err)
		if rerr := tx.Rollback(); rerr != nil {
			log.Println("tx error ::", rerr)
			return "", "", rerr
		}
		return "", "", err
	}

	updated, err := tx.User.
		Update().
		AddAuthTokens(authToken).
		AddAuthTokenIDs(authToken.ID).
		Save(bg)

	log.Println(updated)
	if err != nil {
		log.Println("user update error ::", err)
		if rerr := tx.Rollback(); rerr != nil {
			log.Println("tx error ::", rerr)
			return "", "", rerr
		}
		return "", "", err
	}

	accessSubject := "access_token"
	accessPayload := JSON{
		"user_id": user.ID,
	}

	refreshSubject := "refresh_token"
	refreshPayload := JSON{
		"user_id":  user.ID,
		"token_id": authToken.ID,
	}

	accessToken, _ := GenerateAccessToken(accessPayload, accessSubject)
	refreshToken, _ := GenerateRefreshToken(refreshPayload, refreshSubject)

	return accessToken, refreshToken, tx.Commit()
}

func RefreshUserToken(user *ent.User, client *ent.Client, tokenId, originalRefreshToken string, refreshTokenExp int64) (string, string) {
	now := time.Now().Unix()
	diff := refreshTokenExp - now

	refreshToken := originalRefreshToken
	accessSubject := "access_token"
	accessPayload := JSON{
		"user_id": user.ID,
	}
	accessToken, _ := GenerateAccessToken(accessPayload, accessSubject)
	if diff < 60*60*24*15 {
		log.Println("refreshing....")
		refreshSubject := "refresh_token"
		refreshPayload := JSON{
			"user_id":  user.ID,
			"token_id": tokenId,
		}

		refreshToken, _ = GenerateRefreshToken(refreshPayload, refreshSubject)
	}

	return accessToken, refreshToken
}

func generateToken(payload JSON, subject string, expire time.Duration) (string, error) {
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
