package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func TrendingPostController(ctx *gin.Context) {}

func ListPostController(ctx *gin.Context) {
	cursor := ctx.Query("cursor")
	limit := ctx.Query("limit")
	username := ctx.Query("username")

	limited, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limited > 100 {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorLimited)
		return
	}

	queryObj := dto.ListPostQuery{
		Cursor:   cursor,
		Limit:    limited,
		Username: username,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.ListPostService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}


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

func GetPostController(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	urlSlug := ctx.Param("url_slug")

	if postId == "" || urlSlug == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetPostService(postId, urlSlug, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
