package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
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
