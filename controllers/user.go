package controllers

import (
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func GetUserProfile (ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetUserProfileService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
