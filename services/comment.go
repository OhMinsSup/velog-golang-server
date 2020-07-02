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

func WriteCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	commentRepository := repository.NewCommentRepository(db)
	if err := commentRepository.CreateComment(body, userId); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"comment": true,
	}, http.StatusOK, nil
}

func EditCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	commentRepository := repository.NewCommentRepository(db)
	if err := commentRepository.UpdateComment(body, userId); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"comment": true,
	}, http.StatusOK, nil
}

func RemoveCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	commentRepository := repository.NewCommentRepository(db)
	if err := commentRepository.DeleteComment(body, userId); err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"comment": true,
	}, http.StatusOK, nil
}

func GetCommentListService(postId string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	commentRepository := repository.NewCommentRepository(db)
	comments, err := commentRepository.CommentList(postId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"comments": comments,
	}, http.StatusOK, nil
}

func GetSubCommentListService(postId, commentId string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	return helpers.JSON{
		"comments": true,
	}, http.StatusOK, nil
}
