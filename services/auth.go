package services

import (
	"github.com/OhMinsSup/story-server/database/models"
	emailService "github.com/OhMinsSup/story-server/helpers/email"
	"github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"os"
	"path/filepath"
	"strings"
)

func SendEmailService(email string, db *gorm.DB) (bool, error) {
	exists := false
	var user models.User
	if err := db.Where("email = ?", strings.ToLower(email)).First(&user); err == nil {
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
