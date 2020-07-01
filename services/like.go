package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

func UnLikePostService(postId string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	var postModel models.Post
	if err := db.Where("id = ?", postId).First(&postModel).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFound
	}

	var postLike models.PostLike
	db.Raw(`
		SELECT * FROM "post_likes" AS pl
		WHERE pl.post_id = ? AND pl.user_id = ? ORDER BY pl.id ASC LIMIT 1`, postId, userId).Scan(&postLike)

	if postLike == (models.PostLike{}) {
		return helpers.JSON{
			"liked": postModel.Serialize(),
		}, http.StatusOK, nil
	}

	db.Exec(`DELETE from "post_likes" where id = ?`, postLike.ID)

	var count int
	if err := db.Model(&models.PostLike{}).Where("post_id = ?", postId).Count(&count).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFound
	}

	postModel.Likes = count

	db.Save(&postModel)

	db.Exec(`DELETE from "post_scores" where post_id = ? AND user_id = ? AND type = 'LIKE'`, postId, userId)

	return helpers.JSON{
		"liked": postModel.Serialize(),
	}, http.StatusOK, nil
}

func LikePostService(postId string, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	if userId == "" {
		return nil, http.StatusForbidden, nil
	}

	var postModel models.Post
	if err := db.Where("id = ?", postId).First(&postModel).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFound
	}

	var alreadyLiked models.PostLike
	db.Raw(`
		SELECT * FROM "post_likes" AS pl
		WHERE pl.post_id = ? AND pl.user_id = ? ORDER BY pl.id ASC LIMIT 1`, postId, userId).Scan(&alreadyLiked)

	if alreadyLiked != (models.PostLike{}) {
		return helpers.JSON{
			"liked": postModel.Serialize(),
		}, http.StatusOK, nil
	}

	newPostLike := models.PostLike{
		PostId:    postId,
		UserId:    userId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.NewRecord(newPostLike)
	db.Create(&newPostLike)

	var count int
	if err := db.Model(&models.PostLike{}).Where("post_id = ?", postId).Count(&count).Error; err != nil {
		return nil, http.StatusNotFound, helpers.ErrorNotFound
	}

	postModel.Likes = count

	db.Save(&postModel)

	newPostScore := models.PostScore{
		Type:      "LIKE",
		PostId:    postId,
		UserId:    userId,
		Score:     5,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.NewRecord(newPostScore)
	db.Create(&newPostScore)

	return helpers.JSON{
		"liked": postModel.Serialize(),
	}, http.StatusOK, nil
}
