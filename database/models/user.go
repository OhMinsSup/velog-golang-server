package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID          string `gorm:"primary_key;uuid"`
	Username    string `sql:"index"`
	Email       string `sql:"index"`
	IsCertified bool   `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time  `sql:"index"`
	AuthTokens  []AuthToken `gorm:"foreignkey:UserID"`
	UserProfile UserProfile `gorm:"foreignkey:UserID"`
}

func (user *User) GenerateUserToken(db *gorm.DB) helpers.JSON {
	authToken := AuthToken{
		UserID: user.ID,
	}

	db.NewRecord(authToken)
	db.Create(&authToken)

	accessSubject := "access_token"
	accessPayload := helpers.JSON{
		"user_id": user.ID,
	}

	accessToken, _ := helpers.GenerateAccessToken(accessPayload, accessSubject)

	refreshSubject := "refresh_token"
	refreshPayload := helpers.JSON{
		"user_id":  user.ID,
		"token_id": authToken.ID,
	}

	refreshToken, _ := helpers.GenerateRefreshToken(refreshPayload, refreshSubject)

	return helpers.JSON{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
}

func (user *User) RefreshUserToken() {
}
