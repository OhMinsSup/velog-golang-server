package middlewares

import (
	"errors"
	"fmt"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

// refresh 토큰을 재발급하는 함수
func refresh(db *gorm.DB, ctx *gin.Context, refreshToken string) (string, error) {
	// refresh token 을 decode 를 한다
	decodeTokenData, err := libs.DecodeToken(refreshToken)
	if err != nil {
		return "", errors.New("INVALID_TOKEN")
	}

	payload := decodeTokenData["payload"].(map[string]interface{})

	var user models.User
	// payload 에서 가져온 값이 실제로 존재하는 유저인지 체크
	if err := db.Where("id = ?", payload["user_id"].(string)).First(&user).Error; err != nil {
		return "", errors.New("INVALID_TOKEN")
	}

	userId := payload["user_id"].(string)
	tokenId := payload["token_id"].(string)
	exp := int64(decodeTokenData["exp"].(float64))

	// 토큰값으로 access, refresh 재발급
	tokens := user.RefreshUserToken(tokenId, exp, refreshToken)
	libs.SetCookie(ctx, fmt.Sprintf("%v", tokens["accessToken"]), fmt.Sprintf("%v", tokens["refreshToken"]))
	return userId, nil
}

// ConsumeUser token 검증및 재발급 프로세스
func ConsumeUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.FullPath() == "/auth/logout" {
			context.Next()
			return
		}

		// access token 을 가져온다.
		accessToken, err := context.Cookie("access_token")
		// 못 가져온 경우
		if err != nil {
			// try reading HTTP Header
			authorization := context.Request.Header.Get("Authorization")
			if authorization != "" {
				sp := strings.Split(authorization, "Bearer ")
				// invalid token
				if len(sp) < 1 {
					context.Next()
					return
				}
				// 헤더에 access token이 존재하는 경우에 access token에 값을 넣어준다
				accessToken = sp[1]
			}
		}

		// refresh token 을 가져온다
		refreshToken, err := context.Cookie("refresh_token")
		if err != nil {
			context.Next()
			return
		}

		// access Token refresh token의 값이 없는 경우에는
		if accessToken == "" {
			// invalid token! try token refresh...
			// refresh token이 없는 경우에 다음 미들웨어로 이동
			if refreshToken == "" {
				context.Next()
				return
			}
			// 토큰이 존재하는 경우 다시 token을 재발급 받는다.
			// 그리고 userid값을 받아서 context에 할당
			userId, _ := refresh(db, context, refreshToken)
			context.Set("id", userId)
			context.Next()
			return
		}

		// access token 이 존재하는 경우 token 을 decoed 를 한다
		decodeTokenData, err := libs.DecodeToken(accessToken)
		if err != nil {
			context.Next()
			return
		}

		payload := decodeTokenData["payload"].(map[string]interface{})
		tokenExpire := int64(decodeTokenData["exp"].(float64))
		now := time.Now().Unix()
		diff := tokenExpire - now

		// 만료 시간을 넘은경우 & refreshToken 이 존재하는 경우
		if diff < 60*60 && refreshToken != "" {
			log.Println("refresh...")
			userId, err := refresh(db, context, refreshToken)
			if err != nil {
				context.Next()
				return
			}

			context.Set("id", userId)
			context.Next()
			return
		}

		context.Set("id", payload["user_id"])
		context.Next()
		return
	}
}
