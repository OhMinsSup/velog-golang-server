package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/models"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"
)

func ReadingPostsService(queryObj dto.PostsQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	return helpers.JSON{
		"readings": true,
	}, http.StatusOK, nil
}

func LikePostsService(queryObj dto.PostsQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	log.Println("user", userId)
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	queryCursor := ""
	if queryObj.Cursor != "" {
		var postLike models.PostLike
		if err := db.Where("user_id = ? AND post_id = ?", userId, queryObj.Cursor).First(&postLike).Error; err != nil {
			return nil, http.StatusNotFound, err
		}

		queryCursor = fmt.Sprintf(`AND post_likes.updated_at < '%v' AND post_likes.id != '%v'`, postLike.CreatedAt.Format(time.RFC3339Nano), postLike.ID)
	}

	var data []models.PostLike
	if err := db.Raw(fmt.Sprintf(`
		SELECT * FROM "post_likes"
		INNER JOIN posts ON post_likes.post_id = posts.id
		WHERE post_likes.user_id = '%v'
		%v
		ORDER BY post_likes.id ASC, post_likes.updated_at DESC
		LIMIT %v`, userId, queryCursor, queryObj.Limit)).Scan(&data).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	return helpers.JSON{
		"posts": data,
	}, http.StatusOK, nil
}

func TrendingPostsService(queryObj dto.TrendingPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	var trendingPosts []struct {
		ID    string  `json:"id"`
		Score float64 `json:"score"`
	}
	if err := db.Raw(`
		SELECT posts.id, posts.title, SUM(score) AS score  FROM post_scores
		INNER JOIN posts ON post_scores.post_id = posts.id
		WHERE post_scores.created_at::TIME > now()::TIME - INTERVAL '14 days'::TIME
		AND posts.created_at::TIME > now()::TIME - INTERVAL '3 months'::TIME
		GROUP BY posts.id
		ORDER BY score, posts.id DESC
		OFFSET ?
		LIMIT ?
	`, queryObj.Offset, queryObj.Limit).Scan(&trendingPosts).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if len(trendingPosts) == 0 {
		var empty []struct{}
		return helpers.JSON{
			"ordered": empty,
		}, http.StatusOK, nil
	}

	var ids []string
	for _, postData := range trendingPosts {
		ids = append(ids, postData.ID)
	}

	var ordered []dto.PostRawQueryUserProfileResult
	if err := db.Raw(`
		SELECT p.*, u.id, u.username, u.email, up.display_name, up.thumbnail as user_thumbnail FROM "posts" AS p
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		WHERE p.id IN (?)`, ids).Scan(&ordered).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	return helpers.JSON{
		"ordered": ordered,
	}, http.StatusOK, nil
}

func ListPostsService(body dto.ListPostQuery, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	postRepository := repository.NewPostRepository(db)
	posts, err := postRepository.ListPost(fmt.Sprintf("%v", ctx.MustGet("id")), body)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return helpers.JSON{
		"posts": posts,
	}, http.StatusOK, nil
}
