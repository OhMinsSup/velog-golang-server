package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func WriteCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var postData dto.PostRawQueryUserProfileResult
	if err := db.Raw(`
		SELECT p.*, u.id, u.username, u.email, up.display_name FROM "posts" AS p
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		WHERE p.id = ?`, body.PostId).Scan(&postData).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	var comment models.Comment

	if body.CommentId != "" {
		var commentTarget models.Comment
		if err := db.Where("id = ?", body.CommentId).First(&commentTarget).Error; err != nil {
			return nil, http.StatusNotFound, err
		}

		comment.Level = commentTarget.Level + 1
		comment.ReplyTo = body.CommentId

		if commentTarget.Level > 3 {
			return nil, http.StatusBadRequest, helpers.ErrorMaxCommentLevel
		}

		commentTarget.HasReplies = true
		if err := db.Model(&commentTarget).Update(map[string]interface{}{
			"has_replies": true,
		}).Error; err != nil {
			return nil, http.StatusNotModified, err
		}
	}

	comment.Text = body.Text
	comment.PostId = body.PostId
	comment.UserId = userId

	db.NewRecord(comment)
	db.Create(&comment)

	newPostScore := models.PostScore{
		Type:   "COMMENT",
		PostId: body.PostId,
		UserId: userId,
		Score:  2.5,
	}

	db.NewRecord(newPostScore)
	db.Create(&newPostScore)

	return helpers.JSON{
		"comment": comment.Serialize(),
	}, http.StatusOK, nil
}

func EditCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var comment models.Comment
	if err := db.Where("id = ?", body.CommentId).First(&comment).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if userId != comment.UserId {
		return nil, http.StatusUnauthorized, helpers.ErrorPermission
	}

	if err := db.Model(&comment).Update(map[string]interface{}{
		"text": body.Text,
	}).Error; err != nil {
		return nil, http.StatusNotModified, err
	}

	return helpers.JSON{
		"comment": comment.Serialize(),
	}, http.StatusOK, nil
}

func RemoveCommentService(body dto.CommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var comment models.Comment
	if err := db.Where("id = ?", body.CommentId).First(&comment).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if userId != comment.UserId {
		return nil, http.StatusUnauthorized, helpers.ErrorPermission
	}

	var postData dto.PostRawQueryUserProfileResult
	if err := db.Raw(`
		SELECT p.*, u.id, u.username, u.email, up.display_name FROM "posts" AS p
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		WHERE p.id = ?`, body.PostId).Scan(&postData).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	if err := db.Model(&comment).Update(map[string]interface{}{
		"deleted": true,
	}).Error; err != nil {
		return nil, http.StatusNotModified, err
	}

	var score models.PostScore
	if err := db.Raw(`
		SELECT * FROM "post_scores" AS ps 
		WHERE ps.post_id = ?
		AND ps.user_id = ?
		AND ps.type = 'COMMENT'
		ORDER BY ps.created_at DESC
	`, body.PostId, userId).Scan(&score).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	db.Delete(&score)

	return helpers.JSON{
		"comment": true,
	}, http.StatusOK, nil
}
