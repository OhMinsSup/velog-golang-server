package services

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	emailService "github.com/OhMinsSup/story-server/helpers/email"
	"github.com/OhMinsSup/story-server/models"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

	userData := repository.CreateUserParams{
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

func CodeService(code string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	authRepository := repository.NewAuthRepository(db)

	existsCode, ExistsCode, ExistsErr := authRepository.ExistsCode(code)
	if ExistsErr != nil {
		return nil, ExistsCode, ExistsErr
	}

	if existsCode.Logged {
		return nil, http.StatusForbidden, helpers.ErrorTokenAlreadyUse
	}

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

		registerToken, err := helpers.GenerateRegisterToken(payload, subject)
		if err != nil {
			return nil, http.StatusConflict, err
		}

		return helpers.JSON{
			"email":          existsCode.Email,
			"registerToken": registerToken,
		}, http.StatusOK, nil
	}

	var userProfile models.UserProfile
	if err := db.Where("user_id = ?", user.ID).First(&userProfile).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorUserProfileDefine
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

func SendEmailService(email string, db *gorm.DB) (bool, int, error) {
	authRepository := repository.NewAuthRepository(db)

	exists, ExistsCode, ExistsErr := authRepository.ExistEmail(email)
	if ExistsErr != nil {
		return false, ExistsCode, ExistsErr
	}

	emailAuth, CreateCode, CreateErr := authRepository.CreateEmailAuth(email)
	if CreateErr != nil {
		return false, CreateCode, CreateErr
	}

	wd, err := os.Getwd()
	if err != nil {
		return exists, http.StatusNotFound, err
	}

	var bindData emailService.BindData
	if exists {
		bindData.Keyword = "로그인"
		bindData.Url = "http://localhost:5000/#/email-login?code=" + emailAuth.Code
	} else {
		bindData.Keyword = "회원가입"
		bindData.Url = "http://localhost:5000/#/register?code=" + emailAuth.Code
	}

	addr := os.Getenv("SMTP")
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	host := "smtp.gmail.com"
	port := "587"

	emailConfig := emailService.SetupEmailCredentials(
		host,
		port,
		addr,
		username,
		password,
	)

	sender := emailService.NewEmailSender(&emailConfig, email)

	// create send Email Template
	c := make(chan bool)
	go func() {
		if err := sender.ParseTemplate(filepath.Join(wd, "./statics/emailTemplate.html"), bindData); err != nil {
			c <- true
		}

		if err := sender.Send(email); err != nil {
			c <- true
		}
	}()

	return exists, http.StatusOK, nil
}
