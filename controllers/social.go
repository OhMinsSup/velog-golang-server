package controllers

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/social"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func SocialCallback(ctx *gin.Context) {
	ctx.JSON(200, "HAHA")
}

func SocialRedirect(ctx *gin.Context) {
	provider := ctx.Param("provider")
	next := ctx.Query("next")

	providerType := []string{
		"facebook",
		"github",
		"google",
	}

	if !strings.Contains(strings.Join(providerType, ","), provider) {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorProviderValided)
		return
	}

	loginUrl := social.GenerateSocialLink(provider, next)

	ctx.Redirect(http.StatusMovedPermanently, loginUrl)
}

func GithubCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorNotFound)
		return
	}

	provider := "github"
	accessToken := social.GetGithubAccessToken(code)
	profile := social.GetGithubProfile(accessToken)

	db := ctx.MustGet("db").(*gorm.DB)

	var data models.SocialAccount
	if err := db.Where(&models.SocialAccount{
		SocialId: fmt.Sprint(profile.ID),
		Provider: provider,
	}).First(&data).Error; err != nil {
		panic(err)
		return
	}

	ctx.Set("profile", profile)
	ctx.Set("social", data)
	ctx.Set("accessToken", accessToken)
	ctx.Set("provider", provider)
	ctx.Next()
}

func GoogleCallback(ctx *gin.Context) {
	ctx.Next()
}

func FacebookCallback(ctx *gin.Context) {
	ctx.Next()
}
