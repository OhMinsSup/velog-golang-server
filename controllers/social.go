package controllers

import (
	"github.com/OhMinsSup/story-server/helpers"
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

	ctx.JSON(200,
		helpers.JSON{
			"next":     next,
			"provider": provider,
		})
}
