package middlewares

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

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

		payload := decodeTokenData["payload"].(map[string]interface{})
		tokenExpire := int64(decodeTokenData["exp"].(float64))
		now := time.Now().Unix()
		diff := (tokenExpire * 1000) - now
		log.Println("tokenExpire", tokenExpire)
		log.Println("now", now)
		log.Println("diff", diff)
		log.Println("(tokenExpire * 1000)", tokenExpire*1000)
		log.Println(diff < 1000*60*30)
		refreshToken, errRefresh := context.Cookie("refresh_token")
		if errRefresh != nil {
			context.Next()
			return
		}

		if diff < 1000*60*30 && (refreshToken != "" || errRefresh != nil) {
			log.Println("refreshToken")
		}

		context.Set("id", payload["id"])
		context.Next()
		return
	}
}
