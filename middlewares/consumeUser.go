package middlewares

import (
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

func refresh(db *gorm.DB, ctx *gin.Context, refreshToken string) (string, error) {
	decodeTokenData, errDecode := helpers.DecodeToken(refreshToken)
	if errDecode != nil {
		return "", helpers.ErrorInvalidToken
	}

	payload := decodeTokenData["payload"].(map[string]interface{})

	var user models.User
	if err := db.Where("id = ?", payload["user_id"].(string)).First(&user).Error; err != nil {
		return "", helpers.ErrorInvalidToken
	}

	tokenId := payload["token_id"].(string)
	exp := int64(decodeTokenData["exp"].(float64))

	tokens := user.RefreshUserToken(tokenId, exp, refreshToken)
	ctx.SetCookie("access_token", tokens["accessToken"].(string), 60*60*24, "/", "", false, true)
	ctx.SetCookie("refresh_token", tokens["refreshToken"].(string), 60*60*24*30, "/", "", false, true)
	return payload["token_id"].(string), nil
}

func ConsumeUser(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.FullPath() == "/auth/logout" {
			context.Next()
			return
		}

		accessToken, errAccess := context.Cookie("access_token")
		if errAccess != nil {
			// try reading HTTP Header
			authorization := context.Request.Header.Get("Authorization")
			if authorization == "" {
				context.Next()
				return
			}
			sp := strings.Split(authorization, "Bearer ")
			// invalid token
			if len(sp) < 1 {
				context.Next()
				return
			}
			accessToken = sp[1]
		}

		decodeTokenData, errDecode := helpers.DecodeToken(accessToken)
		if errDecode != nil {
			context.Next()
			return
		}

		refreshToken, errRefresh := context.Cookie("refresh_token")
		if errRefresh != nil {
			context.Next()
			return
		}

		payload := decodeTokenData["payload"].(map[string]interface{})
		tokenExpire := int64(decodeTokenData["exp"].(float64))
		now := time.Now().Unix()
		diff := tokenExpire - now

		if diff < 60*60 && refreshToken != "" {
			log.Println("refreshToken")
			userId, err := refresh(db, context, refreshToken)
			if err != nil {
				panic(err)
				return
			}

			context.Set("id", userId)
			context.Next()
			return
		}

		context.Set("id", payload["id"])
		context.Next()
		return
	}
}
