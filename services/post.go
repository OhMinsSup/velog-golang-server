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
	postId := ctx.Param("post_id")

	if postId == "" {
		return nil, http.StatusBadRequest, nil
	}

	postRepository := repository.NewPostRepository(db)
	post, err := postRepository.GetPost(postId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
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

	isDeleted, err := postRepository.DeletePost(userId, postId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"post": isDeleted,
	}, http.StatusOK, nil
}

func UpdatePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postId := ctx.Param("post_id")

	if postId == "" {
		return nil, http.StatusBadRequest, helpers.ErrorParamRequired
	}

	postRepository := repository.NewPostRepository(db)

	postId, err := postRepository.UpdatePost(body, userId, postId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"post_id": postId,
	}, http.StatusOK, nil
}

func WritePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postRepository := repository.NewPostRepository(db)

	postId, err := postRepository.CreatePost(body, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
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

	if err := postRepository.View(body, userId); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"post": true,
	}, http.StatusOK, nil
}
