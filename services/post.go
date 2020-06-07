package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/fx"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetPostService(postId, urlSlug string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	var postData dto.PostRawQueryResult
	db.Raw(`
		SELECT
		p.*,
		array_agg(t.name) AS tag FROM "posts" AS p
		LEFT OUTER JOIN "posts_tags" AS pt ON pt.post_id = p.id
		LEFT OUTER JOIN "tags" AS t ON t.id = pt.tag_id
		WHERE p.id = ? AND p.url_slug = ?
		GROUP BY p.id, pt.post_id`, postId, urlSlug).Scan(&postData)

	var userData dto.UserRawQueryResult
	db.Raw(`
       SELECT
	   u.*,
	   up.display_name,
       up.short_bio,
	   up.thumbnail
	   FROM "users" AS u
	   LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
	   WHERE u.id = ?`, postData.UserID).Scan(&userData)

	return helpers.JSON{
		"post":       postData,
		"write_user": userData,
	}, http.StatusOK, nil
}

func DeletePostService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postId := ctx.Param("post_id")
	urlSlug := ctx.Param("url_slug")

	if postId == "" || urlSlug == "" {
		return nil, http.StatusBadRequest, helpers.ErrorParamRequired
	}

	var currentPost models.Post
	if err := db.Where("id = ? AND url_slug = ?", postId, urlSlug).First(&currentPost); err != nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFound
	}

	if currentPost.UserID != userId {
		return nil, http.StatusForbidden, helpers.ErrorPermission
	}

	db.Raw("DELETE FROM posts_tags pt WHERE pt.post_id = ?", postId)
	db.Raw("DELETE FROM posts p WHERE p.id = ?", postId)

	return helpers.JSON{
		"post": true,
	}, http.StatusOK, nil
}

func UpdatePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postId := ctx.Param("post_id")
	urlSlug := ctx.Param("url_slug")

	if postId == "" || urlSlug == "" {
		return nil, http.StatusBadRequest, helpers.ErrorParamRequired
	}

	var currentPost models.Post
	if err := db.Where("id = ? AND url_slug = ?", postId, urlSlug).First(&currentPost); err != nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFound
	}

	if currentPost.UserID != userId {
		return nil, http.StatusForbidden, helpers.ErrorPermission
	}

	editPost := models.Post{
		Title:      body.Title,
		Body:       body.Body,
		Thumbnail:  body.Thumbnail,
		IsMarkdown: body.IsMarkdown,
		IsPrivate:  body.IsPrivate,
		UserID:     userId,
	}

	if editPost.IsTemp && !body.IsTemp {
		editPost.ReleasedAt = time.Now()
	}

	processedUrlSlug := helpers.EscapeForUrl(body.UrlSlug)

	if urlSlugDuplicate := db.Where("fk_user_id = ? AND url_slug >= ?", userId, body.UrlSlug).First(&currentPost); urlSlugDuplicate != nil {
		randomString := helpers.GenerateStringName(8)
		processedUrlSlug += "-" + randomString
	}

	editPost.UrlSlug = processedUrlSlug

	db.NewRecord(editPost)
	db.Create(&editPost)

	syncPostTags(body, editPost, db)

	return helpers.JSON{
		"post_id": editPost.ID,
	}, http.StatusOK, nil
}

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

	syncPostTags(body, post, db)

	return helpers.JSON{
		"post_id": post.ID,
	}, http.StatusOK, nil
}

func syncPostTags(body dto.WritePostBody, post models.Post, db *gorm.DB) {
	var tagIds []string
	for iter := 0; iter < len(body.Tag); iter++ {
		currentTag := body.Tag[iter]
		tag := models.TagFindOnCreate(currentTag, db)
		tagIds = append(tagIds, tag.ID)
	}

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
	// 현재 포스트에 등록된 태그 정보
	db.Raw("SELECT DISTINCT pt.tag_id FROM posts p INNER JOIN posts_tags pt ON pt.post_id = p.id WHERE pt.post_id = ?", post.ID).Find(&prevPostTags)

	// get deleted posts_tags Item
	var missing []string
	for _, pt := range prevPostTags {
		if id, prefix := fx.ContainSelector(tagIds, pt.TagId); prefix {
			log.Println("missing", id)
			missing = append(missing, id)
		}
	}

	// get add posts_tags Item
	var adding []string
	for _, t := range tagIds {
		if len(prevPostTags) > 0 {
			for _, pt := range prevPostTags {
				if !strings.Contains(t, pt.TagId) {
					adding = append(adding, t)
				}
			}
		} else {
			adding = append(adding, t)
		}
	}

	// remove tags
	if len(missing) > 0 {
		for _, missingTagId := range missing {
			db.Raw("DELETE FROM posts_tags pt WHERE pt.tag_id = ? AND pt.post_id = ?", missingTagId, post.ID)
		}
	}

	// adding tags
	if len(adding) > 0 {
		for _, addingTagId := range adding {
			postsTags := models.PostsTags{
				PostId: post.ID,
				TagId:  addingTagId,
			}

			db.NewRecord(postsTags)
			db.Create(&postsTags)
		}
	}
}
