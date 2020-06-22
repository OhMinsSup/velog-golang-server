package controllers

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func WriteCommentController(ctx *gin.Context) {
	type WriteCommentBody struct {
		Text string `json:"text"`
	}

	var body WriteCommentBody
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	params := dto.WriteCommentParams{
		PostId: ctx.Param("post_id"),
		Text: body.Text,
	}

	db := ctx.MustGet("db").(*gorm.DB)
	result, code, err := services.WriteCommentService(params, db, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func ReplyWriteCommentController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "")
}
