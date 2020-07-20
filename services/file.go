package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/models"
	"github.com/OhMinsSup/story-server/storage"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func GeneratePresignedUrlService(filename, fileType, refId string, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	db := ctx.MustGet("db").(*gorm.DB)
	sess := ctx.MustGet("sess").(*session.Session)

	bucketName := helpers.GetEnvWithKey("BUCKET_NAME")
	if bucketName == "" {
		return nil, http.StatusBadRequest, helpers.ErrorEnvParamsNotFound
	}

	storageRepository := storage.NewStorageRepository(sess)

	tx := db.Begin()

	var user models.User
	if err := tx.Where("id = ?", userId).Preload("UserProfile").First(&user).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusNotFound, err
	}

	var userImage models.UserImage
	if refId != "" {
		userImage.Path = user.ID + "/" + fmt.Sprintf("%v", time.Now().Unix()) + "/" + fileType + "/" + filename
		userImage.ID = user.ID
		userImage.Type = fileType
	} else {
		userImage.Path = user.ID + "/" + fmt.Sprintf("%v", time.Now().Unix()) + "/" + fileType + "/" + refId + "/" + filename
		userImage.ID = user.ID
		userImage.Type = fileType
		userImage.RefId = refId
	}

	if err := tx.Create(&userImage).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	presignedUrl, err := storageRepository.GetS3PresignedUrl(bucketName, userImage.Path, 15)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return helpers.JSON{
		"presignedUrl": presignedUrl,
	}, http.StatusOK, nil
}
