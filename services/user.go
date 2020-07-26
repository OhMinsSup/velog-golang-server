package services

import (
	"fmt"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetUserProfileService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	username := ctx.Param("username")
	userId := ctx.Query("user_id")

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetUserInfo(username, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"user": user,
	}, 0, nil
}

func GetCurrentUserService(db *gorm.DB, ctx *gin.Context) (helpers.JSON, int, error) {
	userId := fmt.Sprintf("%v", ctx.MustGet("id"))

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.GetCurrentUser(userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return helpers.JSON{
		"user": user,
	}, 0, nil
}
