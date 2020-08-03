package repository

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/models"
	"github.com/jinzhu/gorm"
	"net/http"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetCurrentUser(userId string) (dto.UserRawQueryResult, error) {
	var user dto.UserRawQueryResult
	if err := u.db.Raw(`
       SELECT
	   u.*,
	   up.display_name,
       up.short_bio,
	   up.thumbnail,
 	   um.twitter,
       um.github,
       um.facebook,
       um.email_notification,
       um.email_promotion
	   FROM "users" AS u
	   INNER JOIN "user_profiles" AS up ON up.user_id = u.id
       INNER JOIN "user_meta" As um ON um.user_id = u.id
	   WHERE u.id = ?`, userId).Scan(&user).Error; err != nil {
		return dto.UserRawQueryResult{}, err
	}
	return user, nil
}

func (u *UserRepository) GetUserInfo(username, userId string) (dto.UserRawQueryResult, error) {
	if userId != "" {
		var user dto.UserRawQueryResult
		if err := u.db.Raw(`
		   SELECT
		   u.*,
		   up.display_name,
		   up.short_bio,
		   up.thumbnail
		   FROM "users" AS u
		   INNER JOIN "user_profiles" AS up ON up.user_id = u.id
		   WHERE u.username = ? AND u.id = ?`, username, userId).Scan(&user).Error; err != nil {
			return dto.UserRawQueryResult{}, err
		}
		return user, nil
	}

	var user dto.UserRawQueryResult
	if err := u.db.Raw(`
       SELECT
	   u.*,
	   up.display_name,
       up.short_bio,
	   up.thumbnail,
 	   um.twitter,
       um.github,
       um.facebook,
       um.email_notification,
       um.email_promotion
	   FROM "users" AS u
	   INNER JOIN "user_profiles" AS up ON up.user_id = u.id
       INNER JOIN "user_meta" As um ON um.user_id = u.id
	   WHERE u.username = ?`, username).Scan(&user).Error; err != nil {
		return dto.UserRawQueryResult{}, err
	}
	return user, nil
}

func (u *UserRepository) UpdateProfile(userId, shortBio, displayName, thumbnail string) (int, error) {
	tx := u.db.Begin()

	var userProfile models.UserProfile
	if err := tx.Where("user_id = ?", userId).First(&userProfile).Error; err != nil {
		tx.Rollback()
		return http.StatusNotFound, err
	}

	if err := tx.Model(&userProfile).Updates(models.UserProfile{
		ShortBio:    shortBio,
		DisplayName: displayName,
		Thumbnail:   thumbnail,
	}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (u *UserRepository) UpdateEmailRules(userId string, notification, promotion bool) (int, error) {
	tx := u.db.Begin()

	var userMeta models.UserMeta
	if err := tx.Where("user_id = ?", userId).First(&userMeta).Error; err != nil {
		tx.Rollback()
		return http.StatusNotFound, err
	}

	if err := tx.Model(&userMeta).Updates(models.UserMeta{
		EmailNotification: notification,
		EmailPromotion:    promotion,
	}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (u *UserRepository) UpdateSocialInfo(userId, twitter, facebook, github string) (int, error) {
	tx := u.db.Begin()

	var userMeta models.UserMeta
	if err := tx.Where("user_id = ?", userId).First(&userMeta).Error; err != nil {
		tx.Rollback()
		return http.StatusNotFound, err
	}

	if err := tx.Model(&userMeta).Updates(models.UserMeta{
		Twitter:  twitter,
		Facebook: facebook,
		Github:   github,
	}).Error; err != nil {
		tx.Rollback()
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
