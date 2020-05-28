package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func WritePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	//userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	//processedUrlSlug := helpers.EscapeForUrl(body.UrlSlug)
	//
	//var postModel models.Post
	//if urlSlugDuplicate := db.Where("fk_user_id = ? AND url_slug >= ?", userId, body.UrlSlug).First(&postModel); urlSlugDuplicate != nil {
	//	randomString := helpers.GenerateStringName(8)
	//	processedUrlSlug += "-" + randomString
	//}
	//
	//post := models.Post{
	//	Title:      body.Title,
	//	Body:       body.Body,
	//	Thumbnail:  body.Thumbnail,
	//	UrlSlug:    processedUrlSlug,
	//	IsTemp:     body.IsTemp,
	//	IsMarkdown: body.IsMarkdown,
	//	IsPrivate:  body.IsPrivate,
	//	UserID:     userId,
	//}
	//
	//db.NewRecord(post)
	//db.Create(&post)
	//
	//var tagIds []string
	//for iter := 0; iter < len(body.Tag); iter++ {
	//	currentTag := body.Tag[iter]
	//	tag := models.TagFindOnCreate(currentTag, db)
	//	tagIds = append(tagIds, tag.ID)
	//}
	//
	tagIds := []string{"f2051b8a-a0f8-11ea-ac11-acde48001122", "f2054c0e-a0f8-31sg-b75b-acde48001122", "f2054c0e-w2553-31sg-b75b-acde48001122"}

	// 중복을 제거한 배열을 얻는다.
	var uniqueTagIds []string
	filterTagIds := make(map[string]bool)
	for _, value := range tagIds {
		if _, tagId := filterTagIds[value]; !tagId {
			filterTagIds[value] = true
			uniqueTagIds = append(uniqueTagIds, value)
		}
	}

	type PrevPostTags struct {
		TagId string `json:"tag_id"`
	}
	var prevPostTags []PrevPostTags
	db.Raw("SELECT pt.tag_id FROM posts p INNER JOIN posts_tags pt ON pt.post_id = p.id WHERE pt.post_id = ?", "f204a6be-a0f8-11ea-b131-acde48001122").Find(&prevPostTags)

	var missing []string
	filterMissingTagIds := make(map[string]bool)
	for _, pt := range prevPostTags {
		for _, t := range tagIds {
			if _, f := filterMissingTagIds[pt.TagId]; !f && t != pt.TagId {
				filterMissingTagIds[pt.TagId] = true
				missing = append(missing, pt.TagId)
			}
		}
	}

	fmt.Println("missing", missing)

	return helpers.JSON{
		"post": prevPostTags,
	}, 200, nil
}
