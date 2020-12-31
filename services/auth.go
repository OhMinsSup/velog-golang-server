package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	emailService "github.com/OhMinsSup/story-server/helpers/email"
	"github.com/OhMinsSup/story-server/models"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"time"
)

func SocialRegisterService(body dto.SocialRegisterBody, registerToken string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	authRepository := repository.NewAuthRepository(db)

	decoded, err := helpers.DecodeToken(registerToken)
	if err != nil {
		return nil, http.StatusForbidden, err
	}

	// decoded data (email, id)
	payload := decoded["payload"].(helpers.JSON)
	profile := payload["profile"].(helpers.JSON)

	userData := dto.SocialUserParams{
		Email:       strings.ToLower(payload["email"].(string)),
		Username:    body.UserName,
		UserID:      payload["id"].(string),
		DisplayName: body.DisplayName,
		ShortBio:    body.ShortBio,
		SocialID:    fmt.Sprintf("%c", profile["ID"]),
		AccessToken: fmt.Sprintf("%c", payload["access_token"]),
		Provider:    fmt.Sprintf("%c", payload["provider"]),
	}

	// username, email 이미 존재하는지 체크
	_, existsCode, existsError := authRepository.ExistsByEmailAndUsername(userData.Username, userData.Email)
	if existsError != nil {
		return nil, existsCode, existsError
	}

	user, userProfile, userCode, userError := authRepository.SocialUser(userData)
	if userError != nil {
		return nil, userCode, userError
	}

	tokens := user.GenerateUserToken(db)
	env := helpers.GetEnvWithKey("APP_ENV")
	switch env {
	case "production":
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
		break
	case "development":
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
		break
	default:
		break
	}

	return helpers.JSON{
		"id":           user.ID,
		"username":     user.Username,
		"email":        user.Email,
		"thumbnail":    userProfile.Thumbnail,
		"display_name": userProfile.DisplayName,
		"short_bio":    userProfile.ShortBio,
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
	}, http.StatusOK, nil
}

func LocalRegisterService(body dto.LocalRegisterBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	authRepository := repository.NewAuthRepository(db)
	// email register token deocoded
	decoded, err := helpers.DecodeToken(body.RegisterToken)
	if err != nil {
		return nil, http.StatusForbidden, err
	}

	if decoded["subject"] != "email-register" {
		return nil, http.StatusForbidden, helpers.ErrorInvalidToken
	}

	// decoded data (email, id)
	payload := decoded["payload"].(helpers.JSON)

	userData := dto.CreateUserParams{
		Email:       strings.ToLower(payload["email"].(string)),
		Username:    body.UserName,
		UserID:      payload["id"].(string),
		DisplayName: body.DisplayName,
		ShortBio:    body.ShortBio,
	}

	// username, email 이미 존재하는지 체크
	_, existsCode, existsError := authRepository.ExistsByEmailAndUsername(userData.Username, userData.Email)
	if existsError != nil {
		return nil, existsCode, existsError
	}

	var emailAuth models.EmailAuth
	if existsEmailAuth := db.Where("id = ?", payload["id"].(string)).First(&emailAuth); existsEmailAuth != nil {
		emailAuth.Logged = true
		db.Save(&emailAuth)
	}

	user, userProfile, userCode, userError := authRepository.CreateUser(userData)
	if userError != nil {
		return nil, userCode, userError
	}

	tokens := user.GenerateUserToken(db)
	env := helpers.GetEnvWithKey("APP_ENV")
	switch env {
	case "production":
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
		break
	case "development":
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
		break
	default:
		break
	}

	return helpers.JSON{
		"id":           user.ID,
		"username":     user.Username,
		"email":        user.Email,
		"thumbnail":    userProfile.Thumbnail,
		"display_name": userProfile.DisplayName,
		"short_bio":    userProfile.ShortBio,
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
	}, http.StatusOK, nil
}

// CodeService 이메일 코드로 회원가입을 한 유저의 경우 로그인 아닌 경우 회원가입 분기
func CodeService(code string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	authRepository := repository.NewAuthRepository(db)

	// 코드가 현재 서버에서 발급된 코드인지 확인
	existsCode, statusCode, err := authRepository.ExistsCode(code)
	if err != nil {
		return nil, statusCode, err
	}

	// 이미 인증한 코드의 경우에는 401 error
	if existsCode.Logged {
		return nil, http.StatusForbidden, helpers.ErrorTokenAlreadyUse
	}

	// 발송한 이메일은 발송 후 하루 동안의 유효기간을 가짐 그 이후는 만료처리
	expireTime := existsCode.CreatedAt.AddDate(0, 0, 1).Unix()
	currentTime := time.Now().Unix()
	if currentTime > expireTime || existsCode.Logged {
		return nil, http.StatusForbidden, helpers.ErrorTokenExpiredCode
	}

	// check user with code
	var user models.User
	if err := db.Where("email = ?", strings.ToLower(existsCode.Email)).First(&user).Error; err != nil {
		// 해당 이메일로 등록한 유저가 없는 경우
		subject := "email-register"
		payload := helpers.JSON{
			"email": existsCode.Email,
			"id":    existsCode.ID,
		}

		// 회원가입시 서버에서 발급하는 register token을 가지고 회원가입 절차를 가짐
		registerToken, err := helpers.GenerateRegisterToken(payload, subject)
		if err != nil {
			return nil, http.StatusConflict, err
		}

		return helpers.JSON{
			"email":         existsCode.Email,
			"registerToken": registerToken,
		}, http.StatusOK, nil
	}

	var userProfile models.UserProfile
	if err := db.Where("user_id = ?", user.ID).First(&userProfile).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorUserProfileDefine
	}

	// 토큰 생성
	tokens := user.GenerateUserToken(db)
	env := helpers.GetEnvWithKey("APP_ENV")
	switch env {
	case "production":
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
		break
	case "development":
		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
		break
	default:
		break
	}
	// 해당 이메일로 등록한 유저가 있는 경우
	return helpers.JSON{
		"id":           user.ID,
		"username":     user.Username,
		"email":        user.Email,
		"thumbnail":    userProfile.Thumbnail,
		"display_name": userProfile.DisplayName,
		"short_bio":    userProfile.ShortBio,
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
	}, http.StatusOK, nil
}

// SendEmailService 이메일 로그인및 회원가입을 하기위한 이메일 발송
func SendEmailService(email string, db *gorm.DB) (bool, int, error) {
	authRepository := repository.NewAuthRepository(db)

	// 에메일 검증
	exists, code, err := authRepository.ExistEmail(email)
	if err != nil {
		return false, code, err
	}

	// 이메일 인증 코드 생성
	emailAuth, code, err := authRepository.CreateEmailAuth(email)
	if err != nil {
		return false, code, err
	}

	// 템플릿에 필요한 데이터 바인딩
	var template emailService.AuthTemplate

	template.Subject = "이메일 인증"
	template.Template = "velog-email"
	if exists {
		template.Keyword = "로그인"
		template.Url = "http://127.0.0.1:3000/email-login?code=" + emailAuth.Code
	} else {
		template.Keyword = "회원가입"
		template.Url = "http://127.0.0.1:3000/register?code=" + emailAuth.Code
	}

	// 메일을 생성해서 보낸다
	_, err = emailService.SendTemplateMessage(email, template)
	// 이메일 발송에 실패한 경우
	if err != nil {
		return false, http.StatusInternalServerError, err
	}

	return exists, http.StatusOK, err
}
