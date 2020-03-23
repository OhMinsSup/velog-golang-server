package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorized(context *gin.Context) {
	_, exists := context.Get("user")
	if !exists {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
