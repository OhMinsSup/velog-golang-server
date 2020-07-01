package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func UnLikePostService(postId string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	postRepository := repository.NewPostRepository(db)
	isLiked, err := postRepository.UnLike(postId, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"liked": isLiked,
	}, http.StatusOK, nil
}

func LikePostService(postId string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}


	postRepository := repository.NewPostRepository(db)
	isLiked, err := postRepository.Like(postId, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"liked": isLiked,
	}, http.StatusOK, nil
}
