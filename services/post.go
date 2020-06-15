package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func ListPostService(querys dto.ListPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var posts []models.Post
	query := db.Raw(`
		SELECT * FROM "posts" AS p 
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		ORDER BY p.created_at, p.id DESC
		LIMIT = ?`, querys.Limit)

	if userId == "" {
		query.Where("is_private = false")
	} else {
		query.Where("is_private = false OR post.user_id = ?", userId)
	}

	if querys.Cursor != "" {
		var post models.Post
		if err := db.Where("id = ?", querys.Cursor).Scan(&post).Error; err != nil {
			return nil, http.StatusNotFound, helpers.ErrorInvalidCursor
		}

		query.Where("id = ? AND created_at < ?", post.ID, post.CreatedAt)
		query.Where("created_at = ? AND id < ?", post.CreatedAt, post.ID)
	}

	query.Scan(&posts)
	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}

func GetPostService(postId, urlSlug string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	var postData dto.PostRawQueryResult
	db.Raw(`
		SELECT
		p.*,
		array_agg(t.name) AS tag FROM "posts" AS p
		LEFT OUTER JOIN "posts_tags" AS p  ON pt.post_id = p.id
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

	models.SyncPostTags(body, editPost, db)

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

	models.SyncPostTags(body, post, db)

	return helpers.JSON{
		"post_id": post.ID,
	}, http.StatusOK, nil
}
