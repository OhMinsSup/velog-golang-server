package controllers

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUrlController(ctx *gin.Context) {
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
