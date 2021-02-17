package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func WriteCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	commentRepository := repository.NewCommentRepository(db)
	code, err := commentRepository.CreateComment(body, userId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"comment": true,
	}, http.StatusOK, nil
}

func EditCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	commentRepository := repository.NewCommentRepository(db)
	code, err := commentRepository.UpdateComment(body, userId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"comment": true,
	}, http.StatusOK, nil
}

func RemoveCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	commentRepository := repository.NewCommentRepository(db)
	code, err := commentRepository.DeleteComment(body, userId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"comment": true,
	}, http.StatusOK, nil
}

func GetCommentListService(postId string, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	commentRepository := repository.NewCommentRepository(db)
	comments, code, err := commentRepository.CommentList(postId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"comments": comments,
	}, http.StatusOK, nil
}

func GetSubCommentListService(postId, commentId string, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	commentRepository := repository.NewCommentRepository(db)
	comments, code, err := commentRepository.SubCommentList(commentId, postId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"comments": comments,
	}, http.StatusOK, nil
}
