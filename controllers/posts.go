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

func LikePostsController(ctx *gin.Context) {}

func ReadingPostsController(ctx *gin.Context) {}

func TrendingPostsController(ctx *gin.Context) {
	timeframe := ctx.Query("time")

	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	offset, err := strconv.ParseInt(ctx.Query("offset"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limit > 100 {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorLimited)
		return
	}

	queryObj := dto.TrendingPostQuery{
		Limit:     limit,
		Timeframe: timeframe,
		Offset:    offset,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.TrendingPostsService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func ListPostsController(ctx *gin.Context) {
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
	result, code, err := services.ListPostsService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
