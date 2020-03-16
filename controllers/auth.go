package controllers

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type SendEmailBody struct {
	Email string `json:"email", binding:"exists,email,required"`
}

// SendEmailController Post API email send
func SendEmailController(context *gin.Context) {
	var body SendEmailBody
	if err := context.BindJSON(&body); err != nil {
		context.AbortWithStatus(400)
		return
	}

	db := context.MustGet("db").(*gorm.DB)
	registerd, err := services.SendEmailService(body.Email, db)
	if err != nil {
		return
	}

	context.JSON(200, helpers.JSON{
		"registerd": registerd,
	})
}

func CodeController(ctx *gin.Context) {
	code := ctx.Param("code")

	db := ctx.MustGet("db").(*gorm.DB)

}
