package services

import (
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	emailService "github.com/OhMinsSup/story-server/helpers/email"
	"github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func LocalRegisterService(body dto.LocalRegisterBody, db *gorm.DB) (helpers.JSON, error) {
	decoded, err := helpers.DecodeToken(body.RegisterToken)
	if err != nil {
		return helpers.JSON{}, helpers.ErrorInvalidToken
	}

	if decoded["subject"] != "email-register" {
		return nil, helpers.ErrorInvalidToken
	}

	payload := decoded["payload"].(helpers.JSON)

	var userModel models.User
	if err := db.Where("email = ?", strings.ToLower(payload["email"].(string))).Or("username = ?", body.UserName).First(&userModel).Error; err == nil {
		return nil, helpers.ErrorAlreadyExists
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

	tokens := user.GenerateUserToken(db)

	return helpers.JSON{
		"id":           user.ID,
		"username":     user.Username,
		"email":        user.Email,
		"profile":      userProfile,
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
	}, nil
}

func CodeService(code string, db *gorm.DB) (helpers.JSON, error) {
	var emailAuth models.EmailAuth
	if existsEmail := db.Where("code = ?", code).First(&emailAuth); existsEmail == nil {
		return nil, helpers.ErrorNotFoundEmailAuth
	}

	if emailAuth.Logged {
		return nil, helpers.ErrorTokenAlreadyUse
	}

	expireTime := emailAuth.CreatedAt.AddDate(0, 0, 1).Unix()
	currentTime := time.Now().Unix()
	if currentTime > expireTime || emailAuth.Logged {
		return nil, helpers.ErrorTokenExpiredCode
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
			return helpers.JSON{}, err
		}

		return helpers.JSON{
			"email":          emailAuth.Email,
			"register_token": registerToken,
		}, nil
	}

	var userProfile models.UserProfile
	if err := db.Where("user_id = ?", user.ID).First(&userProfile).Error; err != nil {
		return helpers.JSON{}, helpers.ErrorUserProfileDefine
	}

	tokens := user.GenerateUserToken(db)
	// 해당 이메일로 등록한 유저가 있는 경우
	return helpers.JSON{
		"id":           user.ID,
		"username":     user.Username,
		"email":        user.Email,
		"profile":      userProfile,
		"accessToken":  tokens["accessToken"],
		"refreshToken": tokens["refreshToken"],
	}, nil
}

func SendEmailService(email string, db *gorm.DB) (bool, error) {
	exists := false
	var user models.User
	if existsUser := db.Where("email = ?", strings.ToLower(email)).First(&user); existsUser == nil {
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
		return exists, err
	}

	var bindData emailService.EmailBindData
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

	emailConfig := emailService.SetupEmailCredentials(
		"smtp.gmail.com",
		"587",
		addr,
		username,
		password,
	)

	sender := emailService.NewEmailSender(&emailConfig, email)
	// 해당 이슈를 참고 html 읽는중에 병목현상이 발생
	// https://stackoverflow.com/questions/31361745/slow-performance-of-html-template-in-go-lang-any-workaround
	if err := sender.ParseTemplate(filepath.Join(wd, "./statics/emailTemplate.html"), bindData); err != nil {
		return exists, err
	}

	if err := sender.Send(email); err != nil {
		return exists, err
	}

	return exists, nil
}
