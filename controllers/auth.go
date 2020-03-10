package controllers

import (
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

	var response map[string]interface{}
	response["registerd"] = registerd

	context.JSON(200, response)
}
