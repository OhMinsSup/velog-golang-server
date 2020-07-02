package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
)

func GetTagListController(ctx *gin.Context) {
	cursor := ctx.Query("cursor")
	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	queryObj := dto.TagListQuery{
		Cursor: cursor,
		Limit:  limit,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetTagListService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func TrendingTagListController(ctx *gin.Context) {
	sort := ctx.Query("sort")
	cursor := ctx.Query("cursor")

	limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 64)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sortingType := []string{
		"trending",
		"alphabetical",
	}

	if !strings.Contains(strings.Join(sortingType, ","), sort) {
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorProviderValid)
		return
	}

	queryObj := dto.TagListQuery{
		Cursor: cursor,
		Sort:   sort,
		Limit:  limit,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.TrendingTagListService(queryObj, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
