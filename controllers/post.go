package controllers

import (
	"github.com/OhMinsSup/story-server/app"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// WritePostController - 포스트 작성 API
func WritePostController(ctx *gin.Context) {
	var body dto.WritePostDTO
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, app.BadRequestErrorResponse(err.Error(), nil))
		return
	}

	result, _ := services.WritePostService(body, ctx)
	ctx.JSON(result.Code, result)
}

// UpdatePostController - 포스트 업데이트 API
func UpdatePostController(ctx *gin.Context) {
	var body dto.UpdatePostDTO
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, app.BadRequestErrorResponse(err.Error(), nil))
		return
	}

	result, _ := services.UpdatePostService(body, ctx)
	ctx.JSON(result.Code, result)
}

// ReadPostController - 포스트 상세 API
func ReadPostController(ctx *gin.Context) {
	result, _ := services.ReadPostService(ctx)
	ctx.JSON(result.Code, result)
}

// ListPostController - 포스트 리스트 API
func ListPostController(ctx *gin.Context)  {
	result, _ := services.ListPostService(ctx)
    ctx.JSON(result.Code, result)
}
