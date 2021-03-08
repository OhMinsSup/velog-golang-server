package services

import (
	"context"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/ent"
	socialaccountEnt "github.com/OhMinsSup/story-server/ent/socialaccount"
	userEnt "github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/social"
	"github.com/OhMinsSup/story-server/authorize"
	match "github.com/alexpantyukhin/go-pattern-match"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// SocialRegisterService - 소셜 회원가입 서비스 코드
func SocialRegisterService(ctx *gin.Context, body dto.SocialRegisterDTO) (*app.ResponseException, error) {
	registerToken, err := ctx.Cookie("register_token")
	if err != nil {
		return app.ForbiddenErrorResponse("register authorize is empty or invalid", nil), nil
	}

	decoded, err := authorize.DecodeToken(registerToken)
	if err != nil {
		return app.NotFoundErrorResponse("decoded parsing is missing", nil), nil
	}

	// decoded data (email, id)
	payload := decoded["payload"].(dto.RegisterTokenDTO)

	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	// check duplicates
	exists, err := tx.User.
		Query().
		Where(userEnt.Or(
			userEnt.UsernameEQ(body.UserName),
			userEnt.EmailEQ(payload.Profile.Email),
		)).
		First(bg)

	// 어떤 정보가 이미 존재하는지 에러 메세지를 리턴
	if exists != nil {
		var existMessage string
		if exists.Username == body.UserName {
			existMessage = "username"
		} else {
			existMessage = "email"
		}

		return &app.ResponseException{
			Code:       http.StatusConflict,
			ResultCode: app.ResultErrorCodeAlreadyExists,
			Message:    app.ErrorStatus(existMessage),
			Data:       nil,
		}, nil
	}

	// create user
	user, err := tx.User.
		Create().
		SetUsername(body.UserName).
		SetEmail(payload.Profile.Email).
		SetIsCertified(true).
		Save(bg)

	// 유저 생성이 실패한 경
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	socialAccount, err := tx.SocialAccount.
		Create().
		SetAccessToken(payload.Token).
		SetProvider(payload.Provider).
		SetFkUserID(user.ID).
		SetSocialID(payload.Profile.ID).
		Save(ctx)
	log.Println("ent model socialAccount config", socialAccount)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	userProfile, err := tx.UserProfile.
		Create().
		SetDisplayName(body.DisplayName).
		SetShortBio(body.ShortBio).
		SetUser(user).
		SetUserID(user.ID).
		Save(bg)
	log.Println("ent model user profile", userProfile)

	// 유저 프로필 생성이 실패한 경
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	velogConfig, err := tx.VelogConfig.
		Create().
		SetUser(user).
		SetUserID(user.ID).
		Save(bg)
	log.Println("ent model velog config", velogConfig)

	// velog config 생성이 실패한 경
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
	}

	userMeta, err := tx.UserMeta.
		Create().
		SetUser(user).
		SetUserID(user.ID).
		Save(bg)
	log.Println("ent model user meta", userMeta)

	// user meta 생성이 실패한 경
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return app.TransactionsErrorResponse(rerr.Error(), nil), nil
		}
		return app.InteralServerErrorResponse(err.Error(), nil), nil
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
	accessToken, refreshToken := authorize.GenerateUserToken(user, authToken)
	if accessToken == "" || refreshToken == "" {
		if err := tx.Rollback(); err != nil {
			return app.TransactionsErrorResponse(err.Error(), nil), nil
		}
		return app.InteralServerErrorResponse("authorize is not created", nil), nil
	}

	libs.SetCookie(ctx, accessToken, refreshToken)
	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: libs.JSON{
			"id":           user.ID,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}, tx.Commit()
}

// GetSocialProfileInfoService - cookie에 등록된 registerToken에 등록된 유저 정보를 가져온다.
func GetSocialProfileInfoService(ctx *gin.Context) (*app.ResponseException, error) {
	registerToken, err := ctx.Cookie("register_token")
	if err != nil {
		return app.ForbiddenErrorResponse("register authorize is empty or invalid", nil), nil
	}

	decoded, err := authorize.DecodeToken(registerToken)
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

// SocialCallbackService - 소셜 callback 처리 (accessToken, profile 정보)
func SocialCallbackService(provider string, ctx *gin.Context) (*app.ResponseException, error) {
	code := ctx.Query("code")
	if code == "" {
		return app.BadRequestErrorResponse("CODE IS EMPTY", nil), nil
	}

	result := getSocialInfo(provider, code).(libs.JSON)
	ctx.Set("profile", result["profile"])
	ctx.Set("authorize", result["authorize"])
	ctx.Set("provider", provider)

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data:          nil,
	}, nil
}

// SocialAuthenticationService - 유저 인증 서비스
func SocialAuthenticationService(ctx *gin.Context) (*app.ResponseException, error) {
	token := ctx.MustGet("authorize").(string)
	provider := ctx.MustGet("provider").(string)
	profile := ctx.MustGet("profile").(*social.SocialProfile)

	if profile == nil || token == "" {
		return app.ForbiddenErrorResponse("profile and authorize is empty", nil), nil
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
		accessToken, refreshToken := authorize.GenerateUserToken(user, authToken)
		if accessToken == "" || refreshToken == "" {
			if err := tx.Rollback(); err != nil {
				return app.TransactionsErrorResponse(err.Error(), nil), nil
			}
			return app.InteralServerErrorResponse("authorize is not created", nil), nil
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
		}, tx.Commit()
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
		accessToken, refreshToken := authorize.GenerateUserToken(user, authToken)
		if accessToken == "" || refreshToken == "" {
			if err := tx.Rollback(); err != nil {
				return app.TransactionsErrorResponse(err.Error(), nil), nil
			}
			return app.InteralServerErrorResponse("authorize is not created", nil), nil
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
		}, tx.Commit()
	}

	payload := libs.JSON{
		"profile":     profile,
		"provider":    provider,
		"accessToken": token,
	}

	// 회원가입시 서버에서 발급하는 register authorize 을 가지고 회원가입 절차를 가짐
	registerToken, err := authorize.GenerateRegisterToken(payload, "")
	if registerToken == "" || err != nil {
		return app.ForbiddenErrorResponse("authorize is not created", nil), nil
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
	}, tx.Commit()
}

// getSocialInfo - provider에 따라 authorize, profile 정보를 다르게 가져오는 함수
func getSocialInfo(provider, code string) interface{} {
	_, result := match.Match(provider).
		When("facebook", func() libs.JSON {
			accessToken := social.GetFacebookAccessToken(code)
			profile := social.GetFacebookProfile(accessToken)
			data := libs.JSON{
				"authorize":   accessToken,
				"profile": profile,
			}
			return data
		}).
		When("github", func() libs.JSON {
			accessToken := social.GetGithubAccessToken(code)
			profile := social.GetGithubProfile(accessToken)
			data := libs.JSON{
				"authorize":   accessToken,
				"profile": profile,
			}
			return data
		}).
		When("kakao", func() libs.JSON {
			accessToken := social.GetKakaoAccessToken(code)
			profile := social.GetKakaoProfile(accessToken)
			data := libs.JSON{
				"authorize":   accessToken,
				"profile": profile,
			}
			return data
		}).
		Result()

	return result
}
