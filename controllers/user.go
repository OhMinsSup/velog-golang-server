package controllers

import "github.com/gin-gonic/gin"

func MeController(ctx *gin.Context) {
	ctx.JSON(200, true)
}
