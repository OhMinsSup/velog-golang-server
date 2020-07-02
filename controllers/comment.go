package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetCommentController(ctx *gin.Context) {
	postId := ctx.Param("post_id")

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetCommentListService(postId, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func GetSubCommentController(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	commentId := ctx.Param("comment_id")

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.GetSubCommentListService(postId, commentId, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func WriteCommentController(ctx *gin.Context) {
	type CommentBody struct {
		Text string `json:"text"`
	}

	var body CommentBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	params := dto.CommentParams{
		Text:      body.Text,
		PostId:    ctx.Param("post_id"),
		CommentId: ctx.Query("comment_id"),
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.WriteCommentService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func EditCommentController(ctx *gin.Context) {
	type CommentBody struct {
		Text string `json:"text"`
	}

	var body CommentBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	params := dto.CommentParams{
		Text:      body.Text,
		PostId:    ctx.Param("post_id"),
		CommentId: ctx.Param("comment_id"),
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.EditCommentService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func RemoveCommentController(ctx *gin.Context) {
	params := dto.CommentParams{
		PostId:    ctx.Param("post_id"),
		CommentId: ctx.Param("comment_id"),
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.RemoveCommentService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
