package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

// LocalRegisterController Post API user register
func LocalRegisterController(context *gin.Context) {
	var body dto.LocalRegisterBody
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatus(400)
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	result, err := services.LocalRegisterService(body, db)
	if err != nil {
		log.Println(err)
		return
	}

	context.SetCookie("access_token", result["accessToken"].(string), 60*60*24, "/", "", false, true)
	context.SetCookie("refresh_token", result["refreshToken"].(string), 60*60*24*30, "/", "", false, true)
	context.JSON(200, result)
}

// SendEmailController Post API email send
func SendEmailController(context *gin.Context) {
	var body dto.SendEmailBody
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatus(400)
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	registerd, err := services.SendEmailService(body.Email, db)
	if err != nil {
		log.Println(err)
		return
	}

	context.JSON(200, helpers.JSON{
		"registerd": registerd,
	})
}

// CodeController Get API code exists and login
func CodeController(context *gin.Context) {
	code := context.Param("code")

	db := context.MustGet("db").(*gorm.DB)
	result, err := services.CodeService(code, db)
	if err != nil {
		log.Println(err)
		return
	}
	context.JSON(200, result)
}
