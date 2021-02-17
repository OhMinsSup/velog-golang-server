package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/libs"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func UnLikePostService(postId string, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	postRepository := repository.NewPostRepository(db)
	isLiked, code, err := postRepository.UnLike(postId, userId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"liked": isLiked,
	}, http.StatusOK, nil
}

func LikePostService(postId string, db *gorm.DB, ctx *gin.Context) (libs.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	postRepository := repository.NewPostRepository(db)
	isLiked, code, err := postRepository.Like(postId, userId)
	if err != nil {
		return nil, code, err
	}

	return libs.JSON{
		"liked": isLiked,
	}, http.StatusOK, nil
}
