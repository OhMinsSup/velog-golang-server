package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authorized(context *gin.Context) {
	_, exists := context.Get("id")
	if !exists {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
