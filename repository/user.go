package repository

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/jinzhu/gorm"
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
	   up.thumbnail
	   FROM "users" AS u
	   INNER JOIN "user_profiles" AS up ON up.user_id = u.id
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
	   up.thumbnail
	   FROM "users" AS u
	   INNER JOIN "user_profiles" AS up ON up.user_id = u.id
	   WHERE u.username = ?`, username).Scan(&user).Error; err != nil {
		return dto.UserRawQueryResult{}, err
	}
	return user, nil
}
