package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// CreatePostHistoryController
func CreatePostHistoryController(ctx *gin.Context) {
	var body dto.CreatePostHistoryBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.CreatePostHistoryService(body, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

// UpdatePostController
func UpdatePostController(ctx *gin.Context) {
	var body dto.WritePostBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.UpdatePostService(body, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func DeletePostController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.DeletePostService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

// WritePostController - WritePostController 포스트를 등록하는 API
func WritePostController(ctx *gin.Context) {
	var body dto.WritePostBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)

	result, code, err := services.WritePostService(body, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

// GetPostController - GetPostController 포스트를 가져오는 API
func GetPostController(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetPostService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func PostViewController(ctx *gin.Context) {
	ip := libs.CreateHash(ctx.ClientIP())
	postId := ctx.Param("post_id")

	params := dto.PostViewParams{
		Ip:     ip,
		PostId: postId,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.PostViewService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
