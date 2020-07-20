package controllers

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/services"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
		ctx.AbortWithError(http.StatusBadRequest, helpers.ErrorInvalidData)
		return
	}

	filename := header.Filename
	result, code, err := services.GeneratePresignedUrlService(filename, fileType, refId, ctx)
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	ctx.JSON(code, result)
}

func S3ImageUploadController(ctx *gin.Context) {
	// Source
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	filename := header.Filename
	Bucket := helpers.GetEnvWithKey("BUCKET_NAME")
	Region := helpers.GetEnvWithKey("AWS_REGION")

	sess := ctx.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)
	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(Bucket),
		ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": upload,
		})
		return
	}
	filepath := "https://" + Bucket + "." + "s3-" + Region + ".amazonaws.com/" + filename
	ctx.JSON(http.StatusOK, helpers.JSON{
		"filepath": filepath,
	})
}
