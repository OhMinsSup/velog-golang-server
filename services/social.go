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

func GetSocialProfileInfoService(ctx *gin.Context) (*app.ResponseException, error) {
	registerToken, err := ctx.Cookie("register_token")
	if err != nil {
		return app.ForbiddenErrorResponse("register token is empty or invalid", nil), nil
	}

	decoded, err := libs.DecodeToken(registerToken)
	if err != nil {
		return app.NotFoundErrorResponse("decoded parsing is missing", nil), nil
	}

	return &app.ResponseException{
		Code:          http.StatusMovedPermanently,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"profile": decoded,
		},
	}, nil
}

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

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	socialAccount, err := tx.SocialAccount.Query().Where(
		socialaccountEnt.And(
			socialaccountEnt.SocialIDEQ(profile.ID),
			socialaccountEnt.ProviderEQ(provider),
		),
	).First(bg)

	// socialAccount 정보가 있는 경우
	if !ent.IsNotFound(err) {
		user, err := tx.User.Query().Where(
			userEnt.IDEQ(socialAccount.FkUserID),
		).First(bg)

		if ent.IsNotFound(err) {
			return app.NotFoundErrorResponse("User is missing", nil), nil
		}

		authToken, err := tx.AuthToken.
			Create().
			SetFkUserID(user.ID).
			Save(bg)

		// 토큰 생성이 실패한 경우
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return app.TransactionsErrorResponse(rerr.Error(), nil), nil
			}
			return app.InteralServerErrorResponse(err.Error(), nil), nil
		}

		// 토큰 생성
		accessToken, refreshToken := libs.GenerateUserToken(user, authToken)
		if accessToken == "" || refreshToken == "" {
			if err := tx.Rollback(); err != nil {
				return app.TransactionsErrorResponse(err.Error(), nil), nil
			}
			return app.InteralServerErrorResponse("token is not created", nil), nil
		}

		libs.SetCookie(ctx, accessToken, refreshToken)
		return &app.ResponseException{
			Code:          http.StatusMovedPermanently,
			ResultCode:    0,
			Message:       "",
			ResultMessage: "",
			Data: libs.JSON{
				"redirectUrl": "http://localhost:3000/",
			},
		}, nil
	}

	// Find by email ONLY when email exists
	var user *ent.User
	if profile.Email != "" {
		userInfo, err := tx.User.Query().Where(
			userEnt.EmailEQ(profile.Email),
		).First(bg)

		// 유저가 존재하지 않는 경우 nil 존재하는 경우 UserData
		if ent.IsNotFound(err) {
			user = nil
		} else {
			user = userInfo
		}
	}

	// 유저가 존재하는 경우
	if user != nil {
		authToken, err := tx.AuthToken.
			Create().
			SetFkUserID(user.ID).
			Save(bg)

		// 토큰 생성이 실패한 경우
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return app.TransactionsErrorResponse(rerr.Error(), nil), nil
			}
			return app.InteralServerErrorResponse(err.Error(), nil), nil
		}

		// 토큰 생성
		accessToken, refreshToken := libs.GenerateUserToken(user, authToken)
		if accessToken == "" || refreshToken == "" {
			if err := tx.Rollback(); err != nil {
				return app.TransactionsErrorResponse(err.Error(), nil), nil
			}
			return app.InteralServerErrorResponse("token is not created", nil), nil
		}

		libs.SetCookie(ctx, accessToken, refreshToken)
		return &app.ResponseException{
			Code:          http.StatusMovedPermanently,
			ResultCode:    0,
			Message:       "",
			ResultMessage: "",
			Data: libs.JSON{
				"redirectUrl": "http://localhost:3000/",
			},
		}, nil
	}

	payload := libs.JSON{
		"profile":     profile,
		"provider":    provider,
		"accessToken": token,
	}

	// 회원가입시 서버에서 발급하는 register token 을 가지고 회원가입 절차를 가짐
	registerToken, err := libs.GenerateRegisterToken(payload, "")
	if registerToken == "" || err != nil {
		return app.ForbiddenErrorResponse("token is not created", nil), nil
	}

	libs.SetRegisterCookie(ctx, registerToken)
	return &app.ResponseException{
		Code:          http.StatusMovedPermanently,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"redirectUrl": "http://localhost:3000/register?social=1",
		},
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
