package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/ent"
	emailAuthEnt "github.com/OhMinsSup/story-server/ent/emailauth"
	userEnt "github.com/OhMinsSup/story-server/ent/user"
	userprofileEnt "github.com/OhMinsSup/story-server/ent/userprofile"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/email"
	"github.com/SKAhack/go-shortid"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

func LocalRegisterService(body dto.LocalRegisterBody, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	// email register token deocoded
	decoded, err := helpers.DecodeToken(body.RegisterToken)
	if err != nil {
		return app.ForbiddenErrorResponse(err.Error(), nil), nil
	}

	// velog 서버에서 발행한 정보가 email-register 가 아닌 다른 값인 경우에는 회원가입시 발급한 코드값이 아님
	if decoded["subject"] != "email-register" {
		return app.ForbiddenErrorResponse(errors.New("Not valid token information.").Error(), nil), nil
	}

	// decoded data (email, id)
	payload := decoded["payload"].(helpers.JSON)

	// check duplicates
	exists, err := client.User.
		Query().
		Where(userEnt.Or(
			userEnt.UsernameEQ(body.UserName),
			userEnt.EmailEQ(payload["email"].(string)),
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

	// disable code
	emailAuth, err := client.EmailAuth.
		Update().
		Where(emailAuthEnt.CodeEQ(payload["id"].(string))).
		SetLogged(true).
		Save(bg)

	log.Println(emailAuth)
	if err != nil {
		return app.DBQueryErrorResponse(err.Error(), nil), nil
	}

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	// create user
	user, err := tx.User.
		Create().
		SetUsername(body.UserName).
		SetEmail(payload["email"].(string)).
		SetIsCertified(true).
		Save(bg)

	// 유저 생성이 실패한 경
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return app.TransactionsErrorResponse(err.Error(), nil), nil
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
	if  err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return app.TransactionsErrorResponse(err.Error(), nil), nil
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
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return app.TransactionsErrorResponse(err.Error(), nil), nil
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
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	// 토큰 생성
	accessToken, refreshToken, _ := helpers.GenerateUserToken(user, client, bg)
	if accessToken == "" || refreshToken == "" {
		if err := tx.Rollback(); err != nil {
			return app.TransactionsErrorResponse(err.Error(), nil), nil
		}
		return app.InteralServerErrorResponse("token is not created", nil), nil
	}

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: helpers.JSON{
			"id":           user.ID,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}, tx.Commit()
}

func CodeAuthService(ctx *gin.Context) (*app.ResponseException, error) {
	code := ctx.Param("code")
	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	emailAuth, err := client.EmailAuth.Query().Where(emailAuthEnt.CodeEQ(code)).First(bg)
	if ent.IsNotFound(err) {
		log.Println("emailAuth", err)
		return app.NotFoundErrorResponse(err.Error(), nil), nil
	}
	log.Println("emailAuth", emailAuth)
	if emailAuth.Logged {
		return app.ForbiddenErrorResponse("TOKEN_ALREADY_USE", nil), nil
	}

	// 발송한 이메일은 발송 후 하루 동안의 유효기간을 가짐 그 이후는 만료처리
	expireTime := emailAuth.CreatedAt.AddDate(0, 0, 1).Unix()
	currentTime := time.Now().Unix()
	if currentTime > expireTime || emailAuth.Logged {
		return app.ForbiddenErrorResponse("EXPIRED_CODE", nil), nil
	}

	user, err := client.User.Query().Where(userEnt.EmailEQ(emailAuth.Email)).First(bg)
	if ent.IsNotFound(err) {
		// 해당 이메일로 등록한 유저가 없는 경우
		subject := "email-register"
		payload := helpers.JSON{
			"email": emailAuth.Email,
			"id":    emailAuth.ID,
		}

		// 회원가입시 서버에서 발급하는 register token 을 가지고 회원가입 절차를 가짐
		registerToken, err := helpers.GenerateRegisterToken(payload, subject)
		if err != nil {
			return &app.ResponseException{
				Code:          http.StatusConflict,
				ResultCode:    -1,
				ResultMessage: "",
				Message:       app.GenerateTokenError,
				Data:          nil,
			}, nil
		}

		return &app.ResponseException{
			Code:          http.StatusOK,
			ResultCode:    0,
			Message:       "",
			ResultMessage: "",
			Data: helpers.JSON{
				"email":         emailAuth.Email,
				"registerToken": registerToken,
			},
		}, nil
	}

	userProfile, err := client.UserProfile.
		Query().
		Where(
			userprofileEnt.
				HasUserWith(userEnt.IDEQ(user.ID))).
		First(bg)

	log.Println(userProfile)
	// 토큰 생성
	accessToken, refreshToken, _ := helpers.GenerateUserToken(user, client, bg)
	helpers.SetCookie(ctx, accessToken, refreshToken)

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: helpers.JSON{
			"id":           user.ID,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	}, nil
}

// SendEmailService 이메일 로그인및 회원가입을 하기위한 이메일 발송
func SendEmailService(body dto.SendEmailBody, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)
	bg := context.Background()

	user, _ := client.User.Query().Where(userEnt.EmailEQ(body.Email)).First(bg)

	tx, err := client.Tx(bg)
	if err != nil {
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	// 인증 code Id 값
	shortId := shortid.Generator()

	emailAuth, err := tx.
		EmailAuth.
		Create().
		SetEmail(strings.ToLower(body.Email)).
		SetCode(shortId.Generate()).
		Save(bg)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return app.TransactionsErrorResponse(err.Error(), nil), nil
	}

	var registered bool
	var template email.AuthTemplate
	// 템플릿에 필요한 데이터 바인딩
	template.Subject = "이메일 인증"
	template.Template = "velog-email"
	if user != nil {
		template.Keyword = "로그인"
		template.Url = "http://127.0.0.1:3000/email-login?code=" + emailAuth.Code
		registered = false
	} else {
		template.Keyword = "회원가입"
		template.Url = "http://127.0.0.1:3000/register?code=" + emailAuth.Code
		registered = true
	}

	// 메일을 생성해서 보낸다
	_, err = email.SendTemplateMessage(body.Email, template)
	// 이메일 발송에 실패한 경우
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return app.BadRequestErrorResponse(err.Error(), nil), nil
	}

	return &app.ResponseException{
		Code:          http.StatusOK,
		ResultCode:    0,
		Message:       "",
		ResultMessage: "",
		Data: helpers.JSON{
			"registered": registered,
		},
	}, tx.Commit()
}

//func SocialRegisterService(body dto.SocialRegisterBody, registerToken string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
//	authRepository := repository.NewAuthRepository(db)
//
//	decoded, err := helpers.DecodeToken(registerToken)
//	if err != nil {
//		return nil, http.StatusForbidden, err
//	}
//
//	// decoded data (email, id)
//	payload := decoded["payload"].(helpers.JSON)
//	profile := payload["profile"].(helpers.JSON)
//
//	userData := dto.SocialUserParams{
//		Email:       strings.ToLower(payload["email"].(string)),
//		Username:    body.UserName,
//		UserID:      payload["id"].(string),
//		DisplayName: body.DisplayName,
//		ShortBio:    body.ShortBio,
//		SocialID:    fmt.Sprintf("%c", profile["ID"]),
//		AccessToken: fmt.Sprintf("%c", payload["access_token"]),
//		Provider:    fmt.Sprintf("%c", payload["provider"]),
//	}
//
//	// username, email 이미 존재하는지 체크
//	_, existsCode, existsError := authRepository.ExistsByEmailAndUsername(userData.Username, userData.Email)
//	if existsError != nil {
//		return nil, existsCode, existsError
//	}
//
//	user, userProfile, userCode, userError := authRepository.SocialUser(userData)
//	if userError != nil {
//		return nil, userCode, userError
//	}
//
//	tokens := user.GenerateUserToken(db)
//	env := helpers.GetEnvWithKey("APP_ENV")
//	switch env {
//	case "production":
//		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", ".storeis.vercel.app", true, true)
//		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", ".storeis.vercel.app", true, true)
//		break
//	case "development":
//		ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "localhost", false, true)
//		ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "localhost", false, true)
//		break
//	default:
//		break
//	}
//
//	return helpers.JSON{
//		"id":           user.ID,
//		"username":     user.Username,
//		"email":        user.Email,
//		"thumbnail":    userProfile.Thumbnail,
//		"display_name": userProfile.DisplayName,
//		"short_bio":    userProfile.ShortBio,
//		"accessToken":  tokens["accessToken"],
//		"refreshToken": tokens["refreshToken"],
//	}, http.StatusOK, nil
//}
//
