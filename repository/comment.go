package repository

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/models"
	"github.com/jinzhu/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (c *CommentRepository) CommentCount(postId string) (int64, error) {
	var commentCount int64
	if err := c.db.Model(&models.Comment{}).Where("post_id AND deleted = false", postId).Count(&commentCount).Error; err != nil {
		return 0, err
	}
	return commentCount, nil
}

func (c *CommentRepository) CommentList(postId string) ([]models.Comment, error) {
	var comments []models.Comment
	if err := c.db.Raw(`
		SELECT comments.* FROM "posts"
		LEFT JOIN comments ON comments.post_id = posts.id
		WHERE comments.post_id IN (?)
		AND comments.level = 0
		AND (comments.deleted = FALSE OR comments.has_replies = true)
		ORDER BY comments.created_at ASC`, postId).Scan(&comments).Error; err != nil {
		return nil, nil
	}

	return comments, nil
}

func (c *CommentRepository) SubCommentList(commentId, postId string) ([]models.Comment, error) {
	var comments []models.Comment
	if err := c.db.Where("reply_to = ? AND post_id", commentId, postId).Order("created_at asc").Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *CommentRepository) CreateComment(body dto.CommentParams, userId string) error {
	var postData dto.PostRawQueryUserProfileResult
	if err := c.db.Raw(`
		SELECT p.*, u.id, u.username, u.email, up.display_name, up.thumbnail as user_thumbnail FROM "posts" AS p
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		WHERE p.id = ?`, body.PostId).Scan(&postData).Error; err != nil {
		return err
	}

	tx := c.db.Begin()
	var comment models.Comment

	if body.CommentId != "" {
		var commentTarget models.Comment
		if err := tx.Where("id = ?", body.CommentId).First(&commentTarget).Error; err != nil {
			tx.Rollback()
			return err
		}

		comment.Level = commentTarget.Level + 1
		comment.ReplyTo = body.CommentId

		if commentTarget.Level > 3 {
			tx.Rollback()
			return helpers.ErrorMaxCommentLevel
		}

		commentTarget.HasReplies = true
		if err := tx.Model(&commentTarget).Updates(map[string]interface{}{
			"has_replies": true,
		}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	comment.Text = body.Text
	comment.PostId = body.PostId
	comment.UserId = userId

	if err := tx.Create(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	newPostScore := models.PostScore{
		Type:   "COMMENT",
		PostId: body.PostId,
		UserId: userId,
		Score:  2.5,
	}

	if err := tx.Create(&newPostScore).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *CommentRepository) UpdateComment(body dto.CommentParams, userId string) error {
	tx := c.db.Begin()

	var comment models.Comment
	if err := tx.Where("id = ?", body.CommentId).First(&comment).Error; err != nil {
		tx.Rollback()
		return err
	}

	if userId != comment.UserId {
		tx.Rollback()
		return helpers.ErrorPermission
	}

	if err := tx.Model(&comment).Updates(map[string]interface{}{
		"text": body.Text,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (c *CommentRepository) DeleteComment(body dto.CommentParams, userId string) error {
	var comment models.Comment
	if err := c.db.Where("id = ?", body.CommentId).First(&comment).Error; err != nil {
		return err
	}

	if userId != comment.UserId {
		return helpers.ErrorPermission
	}

	var postData dto.PostRawQueryUserProfileResult
	if err := c.db.Raw(`
		SELECT p.*, u.id, u.username, u.email, up.display_name FROM "posts" AS p
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		WHERE p.id = ?`, body.PostId).Scan(&postData).Error; err != nil {
		return err
	}

	tx := c.db.Begin()

	if err := tx.Model(&comment).Updates(map[string]interface{}{
		"deleted": true,
	}).Error; err != nil {
		return err
	}

	var score models.PostScore
	if err := c.db.Raw(`
		SELECT * FROM "post_scores" AS ps 
		WHERE ps.post_id = ?
		AND ps.user_id = ?
		AND ps.type = 'COMMENT'
		ORDER BY ps.created_at DESC
	`, body.PostId, userId).Scan(&score).Error; err != nil {
		return err
	}

	tx.Delete(&score)
	return tx.Commit().Error
}
