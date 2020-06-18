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

func TrendingPostService(queryObj dto.TrendingPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	return helpers.JSON{}, http.StatusOK, nil
}

func ListPostService(queryObj dto.ListPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	id := ctx.MustGet("id")
	userId := fmt.Sprintf("%v", id)

	var queryIsPrivate string
	if userId == "" {
		queryIsPrivate = "WHERE p.is_private = false"
	} else {
		queryIsPrivate = fmt.Sprintf("WHERE (p.is_private = false OR p.user_id = '%v')", userId)
	}

	queryUsername := ""
	if queryObj.Username != "" {
		queryUsername = fmt.Sprintf(`AND u.username = '%v'`, queryObj.Username)
	}

	queryCursor := ""
	if queryObj.Cursor != "" {
		var post models.Post
		if err := db.Where("id = ?", queryObj.Cursor).First(&post).Error; err != nil {
			return nil, http.StatusNotFound, helpers.ErrorInvalidCursor
		}

		queryCursor = fmt.Sprintf(`AND p.created_at > '%v'`, post.CreatedAt.Format(time.RFC3339Nano))
		//queryCursor =
		//fmt.Sprintf(`AND p.created_at > '%v' OR p.created_at = '%v' AND p.id = '%v'`,
		//post.CreatedAt.Format(time.RFC3339Nano), post.CreatedAt.Format(time.RFC3339Nano), post.ID)
	}

	var posts []dto.PostsRawQueryResult
	query := db.Raw(fmt.Sprintf(`
		SELECT p.*, u.email, u.username, up.display_name, up.short_bio, up.thumbnail AS user_thumbnail FROM "posts" AS p 
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		%v
		%v
		%v
		ORDER BY p.created_at, p.id DESC
		LIMIT ?`, queryIsPrivate, queryUsername, queryCursor), queryObj.Limit)

	query.Scan(&posts)
	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}

func GetPostService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postId := ctx.Param("post_id")
	urlSlug := ctx.Param("url_slug")

	if postId == "" || urlSlug == "" {
		return nil, http.StatusBadRequest, nil
	}

	var postData dto.PostRawQueryResult
	db.Raw(`
		SELECT
		p.*,
		array_agg(t.name) AS tag FROM "posts" AS p
		LEFT OUTER JOIN "posts_tags" AS pt  ON pt.post_id = p.id
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
