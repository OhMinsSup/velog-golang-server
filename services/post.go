package services

import (
	"errors"
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/models"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetPostService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postRepository := repository.NewPostRepository(db)

	post, code, err := postRepository.GetPost(ctx.Param("post_id"))
	if err != nil {
		return nil, code, err
	}

	return post, http.StatusOK, nil
}

func DeletePostService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postId := ctx.Param("post_id")

	if postId == "" {
		return nil, http.StatusBadRequest, helpers.ErrorParamRequired
	}

	postRepository := repository.NewPostRepository(db)

	isDeleted, code, err := postRepository.DeletePost(userId, postId)
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"post": isDeleted,
	}, http.StatusOK, nil
}

// UpdatePostService - 포스트 수정 서비스 코드
func UpdatePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postRepository := repository.NewPostRepository(db)

	postId, code, err := postRepository.UpdatePost(body, fmt.Sprintf("%v", ctx.MustGet("id")), ctx.Param("post_id"))
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"post_id": postId,
	}, http.StatusOK, nil
}

// WritePostService - 포스트 등록 서비스 코드
func WritePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postRepository := repository.NewPostRepository(db)

	postId, code, err := postRepository.CreatePost(body, fmt.Sprintf("%v", ctx.MustGet("id")))
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"post_id": postId,
	}, http.StatusOK, nil
}

// CreatePostHistoryService - 임시 저장에 대한 포스트 히스토리를 생성
func CreatePostHistoryService(body dto.CreatePostHistoryBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	tx := db.Begin()
	postId := ctx.Param("post_id")

	var post models.Post
	if err := tx.Where("id = ?", postId).Find(&post).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId != post.UserID {
		return nil, http.StatusBadRequest, errors.New("NO_PERMISSION")
	}

	postHistory := models.PostHistory{
		Title:      body.Title,
		Body:       body.Body,
		IsMarkdown: body.IsMarkdown,
		PostId:     post.ID,
	}

	if err := tx.Create(&postHistory).Error; err != nil {
		tx.Rollback()
		return nil, http.StatusInternalServerError, err
	}

	var (
		count        int64
		postHistorys []models.PostHistory
	)

	if err := tx.Where("post_id = ?", post.ID).Find(&postHistorys).Count(&count).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if count > 10 {
		if err := tx.Model(&models.PostHistory{}).Where("post_id = ? AND created_at < ?", post.ID, postHistorys[9].CreatedAt).Error; err != nil {
			return nil, http.StatusNotFound, err
		}
	}

	return nil, http.StatusOK, nil
}

func PostViewService(body dto.PostViewParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	postRepository := repository.NewPostRepository(db)

	if code, err := postRepository.View(body, userId); err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"post": true,
	}, http.StatusOK, nil
}
