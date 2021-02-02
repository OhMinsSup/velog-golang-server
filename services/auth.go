package services

import (
	"context"
	"fmt"
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/ent"
	userEnt "github.com/OhMinsSup/story-server/ent/user"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/email"
	"github.com/SKAhack/go-shortid"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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
//// LocalRegisterService - 유저 회원가입 서비스 로직
//func LocalRegisterService(body dto.LocalRegisterBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
//	authRepository := repository.NewAuthRepository(db)
//	// email register token deocoded
//	decoded, err := helpers.DecodeToken(body.RegisterToken)
//	if err != nil {
//		return nil, http.StatusInternalServerError, err
//	}
//
//	// email-register 가 아닌 다른 값인 경우에는 회원가입시 발급한 코드값이 아님
//	if decoded["subject"] != "email-register" {
//		return nil, http.StatusBadRequest, helpers.ErrorInvalidToken
//	}
//
//	// decoded data (email, id)
//	payload := decoded["payload"].(helpers.JSON)
//
//	userData := dto.LocalRegisterDTO{
//		Email:       strings.ToLower(payload["email"].(string)),
//		Username:    body.UserName,
//		UserID:      payload["id"].(string),
//		DisplayName: body.DisplayName,
//		ShortBio:    body.ShortBio,
//	}
//
//	// username, email 이미 존재하는지 체크
//	_, code, err := authRepository.ExistsByEmailAndUsername(userData.Username, userData.Email)
//	if err != nil {
//		return nil, code, err
//	}
//
//	var emailAuth models.EmailAuth
//	// register token 에서 userId을 emailAuth 모델에 존재하는지 체크 존재하는 경우
//	if exists := db.Where("id = ?", payload["id"].(string)).First(&emailAuth); exists != nil {
//		// 존재하는 경우 Logged 를 변경
//		emailAuth.Logged = true
//		db.Save(&emailAuth)
//	}
//
//	// 유저 생성
//	user, code, err := authRepository.CreateUser(userData)
//	if err != nil {
//		return nil, code, err
//	}
//
//	tokens := user.GenerateUserToken(db)
//	helpers.SetCookie(ctx, tokens["accessToken"].(string), tokens["refreshToken"].(string))
//
//	return helpers.JSON{
//		"id":           user.ID,
//		"accessToken":  tokens["accessToken"],
//		"refreshToken": tokens["refreshToken"],
//	}, http.StatusOK, nil
//}
//
//// CodeService 이메일 코드로 회원가입을 한 유저의 경우 로그인 아닌 경우 회원가입 분기
//func CodeService(code string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
//	authRepository := repository.NewAuthRepository(db)
//
//	// 코드가 현재 서버에서 발급된 코드인지 확인
//	existsCode, statusCode, err := authRepository.ExistsCode(code)
//	if err != nil {
//		return nil, statusCode, err
//	}
//
//	// 이미 인증한 코드의 경우에는 401 error
//	if existsCode.Logged {
//		return nil, http.StatusForbidden, helpers.ErrorTokenAlreadyUse
//	}
//
//	// 발송한 이메일은 발송 후 하루 동안의 유효기간을 가짐 그 이후는 만료처리
//	expireTime := existsCode.CreatedAt.AddDate(0, 0, 1).Unix()
//	currentTime := time.Now().Unix()
//	if currentTime > expireTime || existsCode.Logged {
//		return nil, http.StatusForbidden, helpers.ErrorTokenExpiredCode
//	}
//
//	// check user with code
//	var user models.User
//	if err := db.Where("email = ?", strings.ToLower(existsCode.Email)).First(&user).Error; err != nil {
//		// 해당 이메일로 등록한 유저가 없는 경우
//		subject := "email-register"
//		payload := helpers.JSON{
//			"email": existsCode.Email,
//			"id":    existsCode.ID,
//		}
//
//		// 회원가입시 서버에서 발급하는 register token을 가지고 회원가입 절차를 가짐
//		registerToken, err := helpers.GenerateRegisterToken(payload, subject)
//		if err != nil {
//			return nil, http.StatusConflict, err
//		}
//
//		return helpers.JSON{
//			"email":         existsCode.Email,
//			"registerToken": registerToken,
//		}, http.StatusOK, nil
//	}
//
//	var userProfile models.UserProfile
//	if err := db.Where("user_id = ?", user.ID).First(&userProfile).Error; err != nil {
//		return nil, http.StatusNotFound, helpers.ErrorUserProfileDefine
//	}
//
//	// 토큰 생성
//	tokens := user.GenerateUserToken(db)
//	helpers.SetCookie(ctx, tokens["accessToken"].(string), tokens["refreshToken"].(string))
//
//	// 해당 이메일로 등록한 유저가 있는 경우
//	return helpers.JSON{
//		"id":           user.ID,
//		"accessToken":  tokens["accessToken"],
//		"refreshToken": tokens["refreshToken"],
//	}, http.StatusOK, nil
//}

// SendEmailService 이메일 로그인및 회원가입을 하기위한 이메일 발송
func SendEmailService(body dto.SendEmailBody, ctx *gin.Context) (*app.ResponseException, error) {
	client := ctx.MustGet("client").(*ent.Client)
	context := context.Background()

	user, _ := client.User.Query().Where(userEnt.EmailEQ(body.Email)).First(context)

	tx, err := client.Tx(ctx)
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
		Save(context)

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
