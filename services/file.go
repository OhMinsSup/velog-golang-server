package services

import (
	"errors"
	"fmt"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/aws"
	"github.com/OhMinsSup/story-server/models"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

// GeneratePresignedUrlService 파일을 업로드하면 해당 파일을 aws 업로드 할 수 있는 presigned Url 을 생성해서 넘겨준다
// https://qiita.com/daijinload/items/1b0093bcbef36eb3f32e
// https://eunsu-shin.medium.com/pre-signed-url-%EC%9D%84-%EC%9D%B4%EC%9A%A9%ED%95%98%EC%97%AC-s3-%ED%8C%8C%EC%9D%BC-%EA%B3%B5%EC%9C%A0-fbf9261f64d6
func GeneratePresignedUrlService(filename, fileType, refId string, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	db := ctx.MustGet("db").(*gorm.DB)
	sess := ctx.MustGet("sess").(*session.Session)

	bucketName := helpers.GetEnvWithKey("BUCKET_NAME")
	if bucketName == "" {
		return nil, http.StatusBadRequest, errors.New("BUCKET_NAME_IS_EMPTY")
	}

	storageRepository := aws.NewStorageRepository(sess)

	tx := db.Begin()

	// 해당 유저가 존재하는지 체크
	var user models.User
	if err := tx.Where("id = ?", userId).Preload("UserProfile").First(&user).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusNotFound, err
	}

	// user가 업로드한 이미지 정보를 관리하기 위해서 DB 등록
	var userImage models.UserImage
	if refId != "" {
		userImage.Path = user.ID + "/" + fmt.Sprintf("%v", time.Now().Unix()) + "/" + fileType + filename
		userImage.UserID = user.ID
		userImage.Type = fileType
	} else {
		userImage.Path = user.ID + "/" + fmt.Sprintf("%v", time.Now().Unix()) + "/" + fileType + "/" + refId + filename
		userImage.UserID = user.ID
		userImage.Type = fileType
		userImage.RefId = refId
	}

	if err := tx.Create(&userImage).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	// aws s3 presigned Url 생성
	presignedUrl, err := storageRepository.GetS3PresignedUrl(bucketName, userImage.Path, 15)
	if err != nil {
		tx.Rollback()
		return nil, http.StatusBadRequest, err
	}

	return helpers.JSON{
		"imageUrl":     fmt.Sprintf("https://s3.ap-northeast-2.amazonaws.com/s3.images.story.io/%s", userImage.Path),
		"presignedUrl": presignedUrl,
	}, http.StatusOK, tx.Commit().Error
}
