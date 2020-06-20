package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteCommentController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}

func ReplyWriteCommentController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
