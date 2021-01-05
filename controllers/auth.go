package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"net/http"
)

// LocalRegisterController
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

// SendEmailController
func SendEmailController(ctx *gin.Context) {
	var body dto.SendEmailBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
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

// CodeController
func CodeController(ctx *gin.Context) {
	// validation Code Params
	var params = dto.CodeParams{
		Code: ctx.Param("code"),
	}

	if params.Code == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Code is Required"))
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)

	result, code, err := services.CodeService(params.Code, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

// LogoutController
func LogoutController(ctx *gin.Context) {
	helpers.SetCookie(ctx, "", "")
	ctx.Status(http.StatusNoContent)
}
