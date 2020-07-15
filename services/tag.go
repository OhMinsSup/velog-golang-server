package services

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func TrendingTagListService(body dto.TagListQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	tagRepository := repository.NewTagRepository(db)
	tags, code, err := tagRepository.TrendingTagList(body.Cursor, body.Limit)
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"tags": tags,
	}, http.StatusOK, nil
}

func GetTagListService(body dto.TagListQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	tagRepository := repository.NewTagRepository(db)
	tags, code, err := tagRepository.GetTagList(body.Cursor, body.Limit)
	if err != nil {
		return nil, code, err
	}

	return helpers.JSON{
		"tags": tags,
	}, http.StatusOK, nil
}
