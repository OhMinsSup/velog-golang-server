package services

import (
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/helpers/social"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"net/http"
)

func getSocialInfo(provider, code string) (*social.FacebookProfile, *github.User, string) {
	switch provider {
	case "facebook":
		accessToken := social.GetFacebookAccessToken(code)
		profile := social.GetFacebookProfile(accessToken)

		if profile != nil {
			return nil, nil, ""
		}

		return profile, nil, accessToken
	case "github":
		accessToken := social.GetGithubAccessToken(code)
		profile := social.GetGithubProfile(accessToken)

		if profile != nil {
			return nil, nil, ""
		}

		return nil, profile, accessToken
	case "kakao":
		return nil, nil, ""
	case "google":
		return nil, nil, ""
	default:
		return nil, nil, ""
	}
}

func SocialCallbackService(ctx *gin.Context) (*app.ResponseException, error) {
	code := ctx.Query("code")
	if code == "" {
		return app.BadRequestErrorResponse("CODE IS EMPTY", nil), nil
	}

	provider := ctx.Query("provider")
	if provider == "" {
		return app.BadRequestErrorResponse("PROVIDER IS EMPTY", nil), nil
	}

	facebookProfile, githubProfile, token := getSocialInfo(provider, code)
	switch provider {
	case "facebook":
		ctx.Set("profile", facebookProfile)
		break
	case "github":
		ctx.Set("profile", githubProfile)
		break
	}

	ctx.Set("token", token)
	ctx.Set("provider", provider)
	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data:          nil,
	}, nil
}
