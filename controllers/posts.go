package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func LikePostsController(ctx *gin.Context) {
	cursor := ctx.Query("cursor")

	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limit > 100 {
		ctx.AbortWithError(http.StatusBadRequest, libs.ErrorLimited)
		return
	}

	queryObj := dto.PostsQuery{
		Cursor: cursor,
		Limit:  limit,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.LikePostsService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func ReadingPostsController(ctx *gin.Context) {
	cursor := ctx.Query("cursor")

	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limit > 100 {
		ctx.AbortWithError(http.StatusBadRequest, libs.ErrorLimited)
		return
	}

	queryObj := dto.PostsQuery{
		Cursor: cursor,
		Limit:  limit,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.ReadingPostsService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

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
		_ = ctx.AbortWithError(http.StatusBadRequest, libs.ErrorLimited)
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
	username := ctx.Query("username")

	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if limit > 100 {
		ctx.AbortWithError(http.StatusBadRequest, libs.ErrorLimited)
		return
	}

	queryObj := dto.ListPostQuery{
		Cursor:   cursor,
		Limit:    limit,
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
