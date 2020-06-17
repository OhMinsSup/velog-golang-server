package services

import (
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	emailService "github.com/OhMinsSup/story-server/helpers/email"
	"github.com/SKAhack/go-shortid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func LocalRegisterService(body dto.LocalRegisterBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	decoded, err := helpers.DecodeToken(body.RegisterToken)
	if err != nil {
		return nil, http.StatusForbidden, helpers.ErrorInvalidToken
	}

	if decoded["subject"] != "email-register" {
		return nil, http.StatusForbidden, helpers.ErrorInvalidToken
	}

	payload := decoded["payload"].(helpers.JSON)

	var userModel models.User
	if err := db.Where("email = ?", strings.ToLower(payload["email"].(string))).Or("username = ?", body.UserName).First(&userModel).Error; err == nil {
		return nil, http.StatusConflict, helpers.ErrorAlreadyExists
	}

	var emailAuthModel models.EmailAuth
	if existsEmailAuth := db.Where("id = ?", payload["id"].(string)).First(&emailAuthModel); existsEmailAuth != nil {
		emailAuthModel.Logged = true
		db.Save(&emailAuthModel)
	}

	user := models.User{
		Email:       strings.ToLower(payload["email"].(string)),
		IsCertified: true,
		Username:    body.UserName,
	}

	db.NewRecord(user)
	db.Create(&user)

	userProfile := models.UserProfile{
		DisplayName: body.DisplayName,
		ShortBio:    body.ShortBio,
		UserID:      user.ID,
	}

	db.NewRecord(userProfile)
	db.Create(&userProfile)

	velogConfig := models.VelogConfig{
		UserID: user.ID,
	}

	db.NewRecord(velogConfig)
	db.Create(&velogConfig)

	userMeta := models.UserMeta{
		UserID: user.ID,
	}

	db.NewRecord(userMeta)
	db.Create(&userMeta)

	tokens := user.GenerateUserToken(db)
	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "", false, true)
	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "", false, true)

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
	var emailAuth models.EmailAuth
	if existsEmail := db.Where("code = ?", code).First(&emailAuth); existsEmail == nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFoundEmailAuth
	}

	if emailAuth.Logged {
		return nil, http.StatusForbidden, helpers.ErrorTokenAlreadyUse
	}

	expireTime := emailAuth.CreatedAt.AddDate(0, 0, 1).Unix()
	currentTime := time.Now().Unix()
	if currentTime > expireTime || emailAuth.Logged {
		return nil, http.StatusForbidden, helpers.ErrorTokenExpiredCode
	}

	// check user with code
	var user models.User
	if err := db.Where("email = ?", strings.ToLower(emailAuth.Email)).First(&user).Error; err != nil {
		// 해당 이메일로 등록한 유저가 없는 경우
		subject := "email-register"
		payload := helpers.JSON{
			"email": emailAuth.Email,
			"id":    emailAuth.ID,
		}

		registerToken, err := helpers.GenerateRegisterToken(payload, subject)
		if err != nil {
			return nil, http.StatusConflict, err
		}

		return helpers.JSON{
			"email":          emailAuth.Email,
			"register_token": registerToken,
		}, http.StatusOK, nil
	}

	var userProfile models.UserProfile
	if err := db.Where("user_id = ?", user.ID).First(&userProfile).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorUserProfileDefine
	}

	tokens := user.GenerateUserToken(db)
	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "", false, true)
	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "", false, true)
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
	exists := false
	var user models.User
	if err := db.Where("email = ?", strings.ToLower(email)).First(&user).Error; err == nil {
		exists = true
	}

	shortId := shortid.Generator()

	emailAuth := models.EmailAuth{
		Email: strings.ToLower(email),
		Code:  shortId.Generate(),
	}

	db.NewRecord(emailAuth)
	db.Create(&emailAuth)

	wd, err := os.Getwd()
	if err != nil {
		return exists, http.StatusConflict, err
	}

	var bindData emailService.BindData
	if exists {
		bindData.Keyword = "로그인"
		bindData.Url = "https://velog.io/email-login?code=" + emailAuth.Code
	} else {
		bindData.Keyword = "회원가입"
		bindData.Url = "https://velog.io/register?code=" + emailAuth.Code
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

	c := make(chan bool)
	go func() {
		if err := sender.ParseTemplate(filepath.Join(wd, "./statics/emailTemplate.html"), bindData); err != nil {
			c <- true
		}
		if err := sender.Send(email); err != nil {
			c <- true
		}
	}()

	select {
	case snapshot := <-c:
		if snapshot {
			return exists, http.StatusConflict, err
		}
	default:
		return exists, http.StatusOK, nil
	}

	return exists, http.StatusOK, nil
}
