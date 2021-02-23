package services

import (
	"context"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/ent"
	socialaccountEnt "github.com/OhMinsSup/story-server/ent/socialaccount"
	userEnt "github.com/OhMinsSup/story-server/ent/user"
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

func SocialAuthenticationService(ctx *gin.Context) (*app.ResponseException, error) {
	token := ctx.MustGet("token").(string)
	provider := ctx.MustGet("provider").(string)
	profile := ctx.MustGet("profile").(*social.SocialProfile)

	if profile == nil || token == "" {
		return app.ForbiddenErrorResponse("profile and token is empty", nil), nil
	}

	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	socialAccount, err := client.SocialAccount.Query().Where(
		socialaccountEnt.Or(
			socialaccountEnt.SocialIDEQ(profile.ID),
			socialaccountEnt.ProviderEQ(provider),
		),
	).First(bg)

	if ent.IsNotFound(err) {
		_, err := client.User.Query().Where(
			userEnt.EmailEQ(profile.Email),
		).First(bg)

		if !ent.IsNotFound(err) {
			return app.DBQueryErrorResponse(err.Error(), nil), nil
		}

		payload := libs.JSON{
			"profile":     profile,
			"provider":    provider,
			"accessToken": token,
		}

		// 회원가입시 서버에서 발급하는 register token 을 가지고 회원가입 절차를 가짐
		registerToken, err := libs.GenerateRegisterToken(payload, "")
		if registerToken == "" || err != nil {
			return app.InteralServerErrorResponse("token is not created", nil), nil
		}

		redirectUrl := libs.SetRegisterCookie(ctx, registerToken)
		if redirectUrl == "" {
			return app.InteralServerErrorResponse("redirectUrl is not Found", nil), nil
		}

		return &app.ResponseException{
			Code:          http.StatusMovedPermanently,
			ResultCode:    0,
			Message:       "",
			ResultMessage: "",
			Data: libs.JSON{
				"redirectUrl": redirectUrl,
			},
		}, nil
	}

	user, err := client.User.Query().Where(
		userEnt.IDEQ(socialAccount.FkUserID),
	).First(bg)

	if err != nil {
		return app.DBQueryErrorResponse(err.Error(), nil), nil
	}

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"id": user.ID,
		},
	}, err
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
