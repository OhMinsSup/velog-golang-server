package services

import (
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/libs/social"
	match "github.com/alexpantyukhin/go-pattern-match"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SocialCallbackService(provider string, ctx *gin.Context) (*app.ResponseException, error) {
	code := ctx.Query("code")
	if code == "" {
		return app.BadRequestErrorResponse("CODE IS EMPTY", nil), nil
	}

	result := getSocialInfo(provider, code).(libs.JSON)
	ctx.Set("profile", result["profile"])
	ctx.Set("token", result["token"])
	ctx.Set("provider", provider)

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data:          nil,
	}, nil
}

func getSocialInfo(provider, code string) interface{} {
	_, result := match.Match(provider).
		When("facebook", func() libs.JSON {
			accessToken := social.GetFacebookAccessToken(code)
			profile := social.GetFacebookProfile(accessToken)
			data := libs.JSON{
				"token":   accessToken,
				"profile": profile,
			}
			return data
		}).
		When("github", func() libs.JSON {
			accessToken := social.GetGithubAccessToken(code)
			profile := social.GetGithubProfile(accessToken)
			data := libs.JSON{
				"token":   accessToken,
				"profile": profile,
			}
			return data
		}).
		When("kakao", func() libs.JSON {
			accessToken := social.GetKakaoAccessToken(code)
			profile := social.GetKakaoProfile(accessToken)
			data := libs.JSON{
				"token":   accessToken,
				"profile": profile,
			}
			return data
		}).
		Result()

	return result
}
