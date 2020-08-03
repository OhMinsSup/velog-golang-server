package controllers

import (
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetCurrentUserController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetCurrentUserService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func GetUserProfileController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetUserProfileService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func UpdateEmailRulesController(ctx *gin.Context) {
	var body dto.UpdateEmailRulesBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	db := ctx.MustGet("db").(*gorm.DB)
	userRepository := repository.NewUserRepository(db)

	code, err := userRepository.UpdateEmailRules(userId, body.EmailNotification, body.EmailPromotion)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, body)
}

func UpdateSocialController(ctx *gin.Context) {
	var body dto.UpdateSocialInfoBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	db := ctx.MustGet("db").(*gorm.DB)
	userRepository := repository.NewUserRepository(db)

	code, err := userRepository.UpdateSocialInfo(userId, body.Twitter, body.Facebook, body.Github)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, body)
}

func UpdateProfileController(ctx *gin.Context) {
	var body dto.UpdateProfileBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	db := ctx.MustGet("db").(*gorm.DB)
	userRepository := repository.NewUserRepository(db)

	code, err := userRepository.UpdateProfile(userId, body.ShortBio, body.DisplayName, body.Thumbnail)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, body)
}
