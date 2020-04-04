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
func LocalRegisterController(ctx *gin.Context) {
	var body dto.LocalRegisterBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.LocalRegisterService(body, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

// SendEmailController Post API email send
func SendEmailController(ctx *gin.Context) {
	var body dto.SendEmailBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	registerd, code, err := services.SendEmailService(body.Email, db)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, helpers.JSON{
		"registerd": registerd,
	})
}

// CodeController Get API code exists and login
func CodeController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.CodeService(ctx.Param("code"), db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}
	ctx.JSON(code, result)
}
