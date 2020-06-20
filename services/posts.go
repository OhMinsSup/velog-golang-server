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

func ListPostService(body dto.ListPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	var queryIsPrivate string
	if userId == "" {
		queryIsPrivate = "WHERE p.is_private = false"
	} else {
		queryIsPrivate = fmt.Sprintf("WHERE (p.is_private = false OR p.user_id = '%v')", userId)
	}

	queryUsername := ""
	if body.Username != "" {
		queryUsername = fmt.Sprintf(`AND u.username = '%v'`, body.Username)
	}

	queryCursor := ""
	if body.Cursor != "" {
		var post models.Post
		if err := db.Where("id = ?", body.Cursor).First(&post).Error; err != nil {
			return nil, http.StatusNotFound, err
		}

		queryCursor = fmt.Sprintf(`AND p.created_at < '%v'`, post.CreatedAt.Format(time.RFC3339Nano))
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
		ORDER BY p.created_at DESC
		LIMIT ?`, queryIsPrivate, queryUsername, queryCursor), body.Limit)

	query.Scan(&posts)
	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}

func GetPostService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postId := ctx.Param("post_id")

	if postId == "" {
		return nil, http.StatusBadRequest, nil
	}

	var postData dto.PostRawQueryResult
	if err := db.Raw(`
		SELECT
		p.*,
		array_agg(t.name) AS tag FROM "posts" AS p
		LEFT OUTER JOIN "posts_tags" AS pt  ON pt.post_id = p.id
		LEFT OUTER JOIN "tags" AS t ON t.id = pt.tag_id
		WHERE p.id = ?
		GROUP BY p.id, pt.post_id`, postId).Scan(&postData).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	var userData dto.UserRawQueryResult
	if err := db.Raw(`
       SELECT
	   u.*,
	   up.display_name,
       up.short_bio,
	   up.thumbnail
	   FROM "users" AS u
	   LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
	   WHERE u.id = ?`, postData.UserID).Scan(&userData).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	return helpers.JSON{
		"post":       postData,
		"write_user": userData,
	}, http.StatusOK, nil
}

func DeletePostService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postId := ctx.Param("post_id")

	if postId == "" {
		return nil, http.StatusBadRequest, helpers.ErrorParamRequired
	}

	var currentPost models.Post
	if err := db.Where("id = ?", postId).First(&currentPost).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if currentPost.UserID != userId {
		return nil, http.StatusForbidden, helpers.ErrorPermission
	}

	db.Exec("DELETE FROM posts_tags pt WHERE pt.post_id = ?", postId)
	db.Exec("DELETE FROM posts p WHERE p.id = ?", postId)

	return helpers.JSON{
		"post": true,
	}, http.StatusOK, nil
}

func UpdatePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	postId := ctx.Param("post_id")

	if postId == "" {
		return nil, http.StatusBadRequest, helpers.ErrorParamRequired
	}

	var currentPost models.Post
	if err := db.Where("id = ?", postId).First(&currentPost).Error; err != nil {
		return nil, http.StatusNotFound, err
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

	db.Model(&currentPost).Updates(editPost)

	if len(body.Tag) > 0 {
		models.SyncPostTags(body, currentPost, db)
	}

	return helpers.JSON{
		"post_id": currentPost.ID,
	}, http.StatusOK, nil
}

func WritePostService(body dto.WritePostBody, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	post := models.Post{
		Title:      body.Title,
		Body:       body.Body,
		Thumbnail:  body.Thumbnail,
		IsTemp:     body.IsTemp,
		IsMarkdown: body.IsMarkdown,
		IsPrivate:  body.IsPrivate,
		UserID:     userId,
	}

	db.NewRecord(post)
	db.Create(&post)

	if len(body.Tag) > 0 {
		models.SyncPostTags(body, post, db)
	}

	return helpers.JSON{
		"post_id": post.ID,
	}, http.StatusOK, nil
}

func PostViewService(body dto.PostViewParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	var currentRead models.PostRead
	if err := db.Where(`ip_hash = ? AND post_id = ?`, body.Ip, body.PostId).First(&currentRead).Error; err == nil {
		if currentRead == (models.PostRead{}) {
			return helpers.JSON{
				"post": false,
			}, http.StatusOK, nil
		}
	}

	postRead := models.PostRead{
		PostId:    body.PostId,
		UserId:    userId,
		IpHash:    body.Ip,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.NewRecord(postRead)
	db.Create(&postRead)

	var updatePost models.Post
	if err := db.Where("id = ?", body.PostId).First(&updatePost).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if err := db.Model(&updatePost).Update(map[string]interface{}{"views": updatePost.Views + 1, "updated_at": time.Now()}).Error; err != nil {
		return nil, http.StatusNotModified, err
	}

	newPostScore := models.PostScore{
		Type:   "READ",
		PostId: body.PostId,
		UserId: userId,
		Score:  1.0,
	}

	db.NewRecord(newPostScore)
	db.Create(&newPostScore)

	return helpers.JSON{
		"post": true,
	}, http.StatusOK, nil
}
