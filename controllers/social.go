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
	token := ctx.MustGet("token").(string)
	provider := ctx.MustGet("provider").(string)
	profile := ctx.MustGet("profile").(libs.JSON)
	ctx.JSON(200, libs.JSON{
		"token":    token,
		"provider": provider,
		"profile":  profile,
	})
}

func GithubSocialCallback(ctx *gin.Context) {
	//profile := ctx.MustGet("profile").(*github.User)
	//isSocial := ctx.MustGet("isSocial").(bool)
	//accessToken := ctx.MustGet("accessToken").(string)
	//provider := ctx.MustGet("provider").(string)
	//
	//if profile == nil || accessToken == "" {
	//	ctx.AbortWithError(http.StatusForbidden, libs.ErrorForbidden)
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//var user models.User
	//if isSocial {
	//	socialData := ctx.MustGet("social").(*models.SocialAccount)
	//	if err := db.Where("user_id = ?", socialData.ID).First(&user).Error; err != nil {
	//		ctx.AbortWithError(http.StatusNotFound, libs.ErrorUserIsMissing)
	//		return
	//	}
	//
	//	redirectUrl := ""
	//	next := ""
	//	tokens := user.GenerateUserToken(db)
	//	env := libs.GetEnvWithKey("APP_ENV")
	//	switch env {
	//	case "production":
	//		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
	//		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
	//		redirectUrl = "https://storeis.vercel.app/"
	//		break
	//	case "development":
	//		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
	//		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
	//		redirectUrl = "http://localhost:5000/"
	//		break
	//	default:
	//		break
	//	}
	//	ctx.Redirect(http.StatusMovedPermanently, redirectUrl+next)
	//	return
	//}
	//
	//if err := db.Where("email = ?", profile.Email).First(&user).Error; err != nil {
	//	payload := libs.JSON{
	//		"profile":     profile,
	//		"provider":    provider,
	//		"accessToken": accessToken,
	//	}
	//
	//	registerToken, err := libs.GenerateRegisterToken(payload, "")
	//	if err != nil {
	//		ctx.AbortWithError(http.StatusConflict, err)
	//		return
	//	}
	//
	//	env := libs.GetEnvWithKey("APP_ENV")
	//	redirectUrl := ""
	//	switch env {
	//	case "production":
	//		ctx.SetCookie("register_token", registerToken, 60*60, "/", ".storeis.vercel.app", true, true)
	//		redirectUrl = "https://storeis.vercel.app/#/register?social=1"
	//		break
	//	case "development":
	//		ctx.SetCookie("register_token", registerToken, 60*60, "/", "localhost", false, true)
	//		redirectUrl = "http://localhost:5000/#/register?social=1"
	//		break
	//	default:
	//		break
	//	}
	//	ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
	//	return
	//}
	//
	//tokens := user.GenerateUserToken(db)
	//redirectUrl := ""
	//env := libs.GetEnvWithKey("APP_ENV")
	//switch env {
	//case "production":
	//	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
	//	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
	//	redirectUrl = "https://storeis.vercel.app/"
	//	break
	//case "development":
	//	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
	//	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
	//	redirectUrl = "http://localhost:5000/"
	//	break
	//default:
	//	break
	//}
	//ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
}

func FacebookSocialCallback(ctx *gin.Context) {
	//profile := ctx.MustGet("profile").(social.FacebookProfile)
	//isSocial := ctx.MustGet("isSocial").(bool)
	//accessToken := ctx.MustGet("accessToken").(string)
	//provider := ctx.MustGet("provider").(string)
	//
	//if accessToken == "" {
	//	ctx.AbortWithError(http.StatusForbidden, libs.ErrorForbidden)
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//var user models.User
	//if isSocial {
	//	socialData := ctx.MustGet("social").(*models.SocialAccount)
	//	if err := db.Where("user_id = ?", socialData.ID).First(&user).Error; err != nil {
	//		ctx.AbortWithError(http.StatusNotFound, libs.ErrorUserIsMissing)
	//		return
	//	}
	//
	//	tokens := user.GenerateUserToken(db)
	//	env := libs.GetEnvWithKey("APP_ENV")
	//	redirectUrl := ""
	//	next := ""
	//	switch env {
	//	case "production":
	//		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
	//		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
	//		redirectUrl = "https://storeis.vercel.app/"
	//		break
	//	case "development":
	//		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
	//		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
	//		redirectUrl = "http://localhost:5000/"
	//		break
	//	default:
	//		break
	//	}
	//	ctx.Redirect(http.StatusMovedPermanently, redirectUrl+next)
	//	return
	//}
	//
	//if err := db.Where("email = ?", profile.Email).First(&user).Error; err != nil {
	//	payload := libs.JSON{
	//		"profile":     profile,
	//		"provider":    provider,
	//		"accessToken": accessToken,
	//	}
	//
	//	registerToken, err := libs.GenerateRegisterToken(payload, "")
	//	if err != nil {
	//		ctx.AbortWithError(http.StatusConflict, err)
	//		return
	//	}
	//
	//
	//	env := libs.GetEnvWithKey("APP_ENV")
	//	redirectUrl := ""
	//	switch env {
	//	case "production":
	//		ctx.SetCookie("register_token", registerToken, 60*60, "/", ".storeis.vercel.app", true, true)
	//		redirectUrl = "https://storeis.vercel.app/#/register?social=1"
	//		break
	//	case "development":
	//		ctx.SetCookie("register_token", registerToken, 60*60, "/", "localhost", false, true)
	//		redirectUrl = "http://localhost:5000/#/register?social=1"
	//		break
	//	default:
	//		break
	//	}
	//
	//	ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
	//	return
	//}
	//
	//tokens := user.GenerateUserToken(db)
	//env := libs.GetEnvWithKey("APP_ENV")
	//redirectUrl := ""
	//switch env {
	//case "production":
	//	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
	//	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
	//	redirectUrl = "https://storeis.vercel.app/"
	//	break
	//case "development":
	//	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
	//	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
	//	redirectUrl = "http://localhost:5000/"
	//	break
	//default:
	//	break
	//}
	//ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
}
