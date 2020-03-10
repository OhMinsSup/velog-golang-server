package services

import (
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/SKAhack/go-shortid"
	"github.com/jinzhu/gorm"
	"strings"
)

func SendEmailService(email string, db *gorm.DB) (bool, error) {
	exists := false
	var user models.User
	result := db.Where("email = ?", strings.ToLower(email)).First(&user)

	if result.Error != nil {
		exists = true
	}

	shortid := shortid.Generator()

	emailAuth := models.EmailAuth{
		Email: strings.ToLower(email),
		Code:  shortid.Generate(),
	}

	db.NewRecord(emailAuth)
	db.Create(&emailAuth)

	templFileds := &struct {
		Form    string
		To      string
		Subject string
		Body    string
	}{}

	templFileds.To = email
	templFileds.Form = "mins5190@gmail.com"

	if exists {
		templFileds.Subject = "Story 로그인"
	} else {
		templFileds.Subject = "Story 회원가입"
	}




	return true, nil
}
