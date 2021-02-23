package libs

import "github.com/gin-gonic/gin"

func SetCookie(ctx *gin.Context, accessToken, refreshToken string) {
	env := GetEnvWithKey("APP_ENV")
	switch env {
	case "production":
		ctx.SetCookie("access_token", accessToken, 60*60*24, "/", ".storeis.vercel.app", true, true)
		ctx.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", ".storeis.vercel.app", true, true)
		break
	case "development":
		ctx.SetCookie("access_token", accessToken, 60*60*24, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", refreshToken, 60*60*24*30, "/", "localhost", false, true)
		break
	default:
		break
	}
}

func SetRegisterCookie(ctx *gin.Context, registerToken string) string {
	env := GetEnvWithKey("APP_ENV")
	switch env {
	case "production":
		ctx.SetCookie("register_token", registerToken, 60*60, "/", ".storeis.vercel.app", true, true)
		return "https://storeis.vercel.app/#/register?social=1"
	case "development":
		ctx.SetCookie("register_token", registerToken, 60*60, "/", "localhost", false, true)
		return "http://localhost:5000/#/register?social=1"
	default:
		return ""
	}
}
