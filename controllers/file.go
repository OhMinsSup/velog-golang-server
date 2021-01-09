package controllers

import (
	"errors"
	"github.com/OhMinsSup/story-server/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GeneratePresignedUrlController
func GeneratePresignedUrlController(ctx *gin.Context) {
	// Source
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	refId := ctx.Request.FormValue("refId")
	fileType := ctx.Request.FormValue("fileType")

	if fileType == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("INVALID_DATA"))
		return
	}

	result, code, err := services.GeneratePresignedUrlService(header.Filename, fileType, refId, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}
