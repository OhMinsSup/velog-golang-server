package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// LocalRegisterController Post API user register
func LocalRegisterController(context *gin.Context) {
	var body dto.LocalRegisterBody
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	result, code, err := services.LocalRegisterService(body, db, context)
	if err != nil {
		context.AbortWithError(code, err)
		return
	}

	context.JSON(code, result)
}

// SendEmailController Post API email send
func SendEmailController(context *gin.Context) {
	var body dto.SendEmailBody
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	registerd, code, err := services.SendEmailService(body.Email, db)
	if err != nil {
		context.AbortWithError(code, err)
		return
	}

	context.JSON(code, helpers.JSON{
		"registerd": registerd,
	})
}

// CodeController Get API code exists and login
func CodeController(context *gin.Context) {
	db := context.MustGet("db").(*gorm.DB)
	result, code, err := services.CodeService(context.Param("code"), db, context)
	if err != nil {
		context.AbortWithError(code, err)
		return
	}
	context.JSON(code, result)
}
