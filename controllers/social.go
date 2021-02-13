package controllers

import (
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/helpers/social"
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

	loginUrl := social.GenerateSocialLink(provider, next)
	ctx.Redirect(http.StatusMovedPermanently, loginUrl)
}

func SocialCallbackController (ctx *gin.Context) {
	result, _  := services.SocialCallbackService(ctx)
	if result.Code != http.StatusOK {
		app.UnAuthorizedErrorResponse("Social Auth Error", nil)
		return
	}

	ctx.Next()
}

func SocialProfileController(ctx *gin.Context) {
	//registerToken, err := ctx.Cookie("register_token")
	//if err != nil {
	//	ctx.AbortWithError(http.StatusUnauthorized, err)
	//	return
	//}
	//
	//decoded, err := helpers.DecodeToken(registerToken)
	//if err != nil {
	//	ctx.AbortWithError(http.StatusForbidden, err)
	//	return
	//}
	//
	//// decoded data (email, id)
	//payload := decoded["payload"].(helpers.JSON)
	//profile := payload["profile"].(helpers.JSON)
	//
	//ctx.JSON(http.StatusOK, profile)
}

func SocialRegisterController(ctx *gin.Context) {
	//registerToken, err := ctx.Cookie("register_token")
	//if err != nil {
	//	ctx.AbortWithStatus(http.StatusUnauthorized)
	//	return
	//}
	//
	//var body dto.SocialRegisterBody
	//if err := ctx.BindJSON(&body); err != nil {
	//	ctx.AbortWithStatus(http.StatusBadRequest)
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//result, code, err := services.SocialRegisterService(body, registerToken, db, ctx)
	//if err != nil {
	//	ctx.AbortWithError(code, err)
	//	return
	//}
	//
	//ctx.JSON(code, result)
}

//func SocialRedirect(ctx *gin.Context) {
//provider := ctx.Param("provider")
//next := ctx.Query("next")
//
//providerType := []string{
//	"facebook",
//	"github",
//	"google",
//}
//
//if !strings.Contains(strings.Join(providerType, ","), provider) {
//	ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorProviderValid)
//	return
//}
//
//loginUrl := social.GenerateSocialLink(provider, next)
//ctx.Redirect(http.StatusMovedPermanently, loginUrl)
//}

func FacebookCallback(ctx *gin.Context) {
	//code := ctx.Query("code")
	//if code == "" {
	//	ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorNotFound)
	//	return
	//}
	//
	//provider := "facebook"
	//accessToken := social.GetFacebookAccessToken(code)
	//profile := social.GetFacebookProfile(accessToken)
	//db := ctx.MustGet("db").(*gorm.DB)
	//
	//var data models.SocialAccount
	//if err := db.Where(&models.SocialAccount{
	//	SocialId: fmt.Sprintf("%v", profile.ID),
	//	Provider: provider,
	//}).First(&data); err != nil {
	//	ctx.Set("social", nil)
	//	ctx.Set("isSocial", false)
	//} else {
	//	ctx.Set("social", data)
	//	ctx.Set("isSocial", true)
	//}
	//
	//ctx.Set("profile", profile)
	//ctx.Set("accessToken", accessToken)
	//ctx.Set("provider", provider)
	//ctx.Next()
	//return
}

func GithubCallback(ctx *gin.Context) {
	//code := ctx.Query("code")
	//if code == "" {
	//	ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorNotFound)
	//	return
	//}
	//
	//provider := "github"
	//accessToken := social.GetGithubAccessToken(code)
	//profile := social.GetGithubProfile(accessToken)
	//db := ctx.MustGet("db").(*gorm.DB)
	//
	//var data models.SocialAccount
	//if err := db.Where(&models.SocialAccount{
	//	SocialId: fmt.Sprintf("%v", profile.ID),
	//	Provider: provider,
	//}).First(&data); err != nil {
	//	ctx.Set("social", nil)
	//	ctx.Set("isSocial", false)
	//} else {
	//	ctx.Set("social", data)
	//	ctx.Set("isSocial", true)
	//}
	//
	//ctx.Set("profile", profile)
	//ctx.Set("accessToken", accessToken)
	//ctx.Set("provider", provider)
	//ctx.Next()
	//return
}

func GithubSocialCallback(ctx *gin.Context) {
	//profile := ctx.MustGet("profile").(*github.User)
	//isSocial := ctx.MustGet("isSocial").(bool)
	//accessToken := ctx.MustGet("accessToken").(string)
	//provider := ctx.MustGet("provider").(string)
	//
	//if profile == nil || accessToken == "" {
	//	ctx.AbortWithError(http.StatusForbidden, helpers.ErrorForbidden)
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//var user models.User
	//if isSocial {
	//	socialData := ctx.MustGet("social").(*models.SocialAccount)
	//	if err := db.Where("user_id = ?", socialData.ID).First(&user).Error; err != nil {
	//		ctx.AbortWithError(http.StatusNotFound, helpers.ErrorUserIsMissing)
	//		return
	//	}
	//
	//	redirectUrl := ""
	//	next := ""
	//	tokens := user.GenerateUserToken(db)
	//	env := helpers.GetEnvWithKey("APP_ENV")
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
	//	payload := helpers.JSON{
	//		"profile":     profile,
	//		"provider":    provider,
	//		"accessToken": accessToken,
	//	}
	//
	//	registerToken, err := helpers.GenerateRegisterToken(payload, "")
	//	if err != nil {
	//		ctx.AbortWithError(http.StatusConflict, err)
	//		return
	//	}
	//
	//	env := helpers.GetEnvWithKey("APP_ENV")
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
	//env := helpers.GetEnvWithKey("APP_ENV")
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
	//	ctx.AbortWithError(http.StatusForbidden, helpers.ErrorForbidden)
	//	return
	//}
	//
	//db := ctx.MustGet("db").(*gorm.DB)
	//var user models.User
	//if isSocial {
	//	socialData := ctx.MustGet("social").(*models.SocialAccount)
	//	if err := db.Where("user_id = ?", socialData.ID).First(&user).Error; err != nil {
	//		ctx.AbortWithError(http.StatusNotFound, helpers.ErrorUserIsMissing)
	//		return
	//	}
	//
	//	tokens := user.GenerateUserToken(db)
	//	env := helpers.GetEnvWithKey("APP_ENV")
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
	//	payload := helpers.JSON{
	//		"profile":     profile,
	//		"provider":    provider,
	//		"accessToken": accessToken,
	//	}
	//
	//	registerToken, err := helpers.GenerateRegisterToken(payload, "")
	//	if err != nil {
	//		ctx.AbortWithError(http.StatusConflict, err)
	//		return
	//	}
	//
	//
	//	env := helpers.GetEnvWithKey("APP_ENV")
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
	//env := helpers.GetEnvWithKey("APP_ENV")
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
