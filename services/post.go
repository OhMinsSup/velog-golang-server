package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func WritePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	processedUrlSlug := helpers.EscapeForUrl(body.UrlSlug)

	var postModel models.Post
	if urlSlugDuplicate := db.Where("fk_user_id = ? AND url_slug >= ?", userId, body.UrlSlug).First(&postModel); urlSlugDuplicate != nil {
		randomString := helpers.GenerateStringName(8)
		processedUrlSlug += "-" + randomString
	}

	post := models.Post{
		Title:      body.Title,
		Body:       body.Body,
		Thumbnail:  body.Thumbnail,
		UrlSlug:    processedUrlSlug,
		IsTemp:     body.IsTemp,
		IsMarkdown: body.IsMarkdown,
		IsPrivate:  body.IsPrivate,
		UserID:     userId,
	}

	db.NewRecord(post)
	db.Create(&post)

	var tagIds []string
	for iter := 0; iter < len(body.Tag); iter++ {
		currentTag := body.Tag[iter]
		tag := models.TagFindOnCreate(currentTag, db)
		tagIds = append(tagIds, tag.ID)
	}

	var uniqueTagIds []string

	// 중복을 제거한 배열을 얻는다.
	filterTagIds := make(map[string]bool)
	for _, value := range tagIds {
		if _, tagId := filterTagIds[value]; !tagId {
			filterTagIds[value] = true
			uniqueTagIds = append(uniqueTagIds, value)
		}
	}

	return helpers.JSON{
		"post": post,
	}, 200, nil
}
