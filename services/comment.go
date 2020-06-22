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

func WriteCommentService(body dto.WriteCommentParams, db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))
	var postData dto.PostRawQueryUserProfileResult
	if err := db.Raw(`
		SELECT p.*, u.id, u.username, u.email, up.display_name FROM "posts" AS p
		LEFT OUTER JOIN "users" AS u ON u.id = p.user_id
		LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
		WHERE p.id = ?`, body.PostId).Scan(&postData).Error; err != nil {
		return nil, http.StatusNotFound, err
	}

	newComment := models.Comment{
		Text:   body.Text,
		PostId: body.PostId,
		UserId: userId,
	}

	db.NewRecord(newComment)
	db.Create(&newComment)

	newPostScore := models.PostScore{
		Type:   "COMMENT",
		PostId: body.PostId,
		UserId: userId,
		Score:  2.5,
	}

	db.NewRecord(newPostScore)
	db.Create(&newPostScore)

	return helpers.JSON{}, http.StatusOK, nil
}

func ReplyWriteCommentService() {}
