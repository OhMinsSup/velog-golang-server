package controllers

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/social"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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

	accessToken := social.GetGithubAccessToken(code)

	ctx.Next()
}
