package controllers

import (
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func LikePostController(ctx *gin.Context) {
	postId := ctx.Param("post_id")

	if postId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.LikePostService(postId, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func UnLikePostController(ctx *gin.Context) {
	postId := ctx.Param("post_id")

	if postId == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.UnLikePostService(postId, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
