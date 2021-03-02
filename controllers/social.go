package controllers

import (
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/libs/social"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SocialRegisterController(ctx *gin.Context) {
	ctx.JSON(200, libs.JSON{
		"ok": true,
	})
}

func SocialRedirectController(ctx *gin.Context) {
	provider := ctx.Param("provider")
	next := ctx.Query("next")

	providerType := []string{
		"facebook",
		"github",
		"kakao",
	}

	if !strings.Contains(strings.Join(providerType, ","), provider) {
		ctx.JSON(http.StatusBadRequest, "PROVIDER_VALID")
		return
	}

	redirectUrl := social.GenerateSocialLink(provider, next)
	ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func SocialFacebookCallbackController(ctx *gin.Context) {
	result, _ := services.SocialCallbackService("facebook", ctx)
	if result.Code != http.StatusOK {
		app.UnAuthorizedErrorResponse("Social Auth Error", nil)
		return
	}
	ctx.Next()
}

func SocialGithubCallbackController(ctx *gin.Context) {
	result, _ := services.SocialCallbackService("github", ctx)
	if result.Code != http.StatusOK {
		app.UnAuthorizedErrorResponse("Social Auth Error", nil)
		return
	}
	ctx.Next()
}

func SocialKakaoCallbackController(ctx *gin.Context) {
	result, _ := services.SocialCallbackService("kakao", ctx)
	if result.Code != http.StatusOK {
		app.UnAuthorizedErrorResponse("Social Auth Error", nil)
		return
	}
	ctx.Next()
}

func SocialCallbackController(ctx *gin.Context) {
	result, _ := services.SocialAuthenticationService(ctx)
	if result.Code == http.StatusMovedPermanently {
		ctx.Redirect(http.StatusMovedPermanently, result.Data["redirectUrl"].(string))
		return
	}
	ctx.JSON(result.Code, result)
}

func GetSocialProfileController(ctx *gin.Context) {
	result, _ := services.GetSocialProfileInfoService(ctx)
	ctx.JSON(result.Code, result)
}
