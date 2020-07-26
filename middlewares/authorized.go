package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorized(ctx *gin.Context) {
	_, exists := ctx.Get("id")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
