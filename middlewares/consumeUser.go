package middlewares

import (
	"fmt"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/models"
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
	accessT := fmt.Sprintf("%v", tokens["accessToken"])
	refreshT := fmt.Sprintf("%v", tokens["refreshToken"])

	env := helpers.GetEnvWithKey("APP_ENV")
	switch env {
	case "production":
		ctx.SetCookie("access_token", accessT, 60*60*24, "/", ".storeis.vercel.app", true, true)
		ctx.SetCookie("refresh_token", refreshT, 60*60*24*30, "/", ".storeis.vercel.app", true, true)
		break
	case "development":
		ctx.SetCookie("access_token", accessT, 60*60*24, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", refreshT, 60*60*24*30, "/", "localhost", false, true)
		break
	default:
		break
	}

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
			if authorization != "" {
				sp := strings.Split(authorization, "Bearer ")
				// invalid token
				if len(sp) < 1 {
					context.Next()
					return
				}
				accessToken = sp[1]
			}
		}

		refreshToken, errRefresh := context.Cookie("refresh_token")
		if errRefresh != nil {
			context.Next()
			return
		}

		decodeTokenData, errDecode := helpers.DecodeToken(accessToken)
		if errDecode != nil {
			userId, err := refresh(db, context, refreshToken)
			if err != nil {
				context.Set("id", "")
				context.Next()
				return
			}

			context.Set("id", userId)
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
				context.Set("id", "")
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
