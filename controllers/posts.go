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

func TrendingPostController(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	timeframe := ctx.Query("time")

	limited, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limited > 100 {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorLimited)
		return
	}

	queryObj := dto.TrendingPostQuery{
		Limit:     limited,
		Timeframe: timeframe,
		Offset:    offset,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.TrendingPostService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

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
	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetPostService(db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func PostViewController(ctx *gin.Context) {
	ip := helpers.CreateHash(ctx.ClientIP())
	postId := ctx.Param("post_id")
	urlSlug := ctx.Param("url_slug")

	params := dto.PostViewParams{
		Ip:      ip,
		PostId:  postId,
		UrlSlug: urlSlug,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.PostViewService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
