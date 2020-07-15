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

func ReadingPostsService(queryObj dto.PostsQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	postRepository := repository.NewPostRepository(db)
	posts, code, err := postRepository.ReadingPostList(userId, queryObj)
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}

func LikePostsService(queryObj dto.PostsQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	postRepository := repository.NewPostRepository(db)
	posts, code, err := postRepository.LikePostList(userId, queryObj)
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}

func TrendingPostsService(queryObj dto.TrendingPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postRepository := repository.NewPostRepository(db)
	posts, code, err := postRepository.TrendingPostList(queryObj)
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}

func ListPostsService(body dto.ListPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postRepository := repository.NewPostRepository(db)
	posts, code, err := postRepository.PostList(fmt.Sprintf("%v", ctx.MustGet("id")), body)
	if err != nil {
		return nil, code, err
	}
	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}
