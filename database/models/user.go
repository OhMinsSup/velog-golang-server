package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	ID          string      `gorm:"primary_key;uuid"json:"id"`
	Username    string      `sql:"index"json:"username"`
	Email       string      `sql:"index"json:"email"`
	IsCertified bool        `gorm:"default:false"json:"is_certified"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	//DeletedAt   *time.Time  `sql:"index"json:"deleted_at"`
	AuthTokens  []AuthToken `gorm:"foreignkey:UserID"json:"auth_tokens"`
	UserProfile UserProfile `gorm:"foreignkey:UserID"json:"user_profile"`
	UserMeta    UserMeta    `gorm:"foreignkey:UserID"json:"user_meta"`
	VelogConfig VelogConfig `gorm:"foreignkey:UserID"json:"velog_config"`
	PostScore   []PostScore `gorm:"polymorphic:Owner;"`
	PostRead    []PostRead  `gorm:"polymorphic:Owner;"`
	PostLike    []PostLike  `gorm:"polymorphic:Owner;"`
}

func (u User) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"email":        u.Email,
		"display_name": u.UserProfile.DisplayName,
		"short_bio":    u.UserProfile.ShortBio,
		"thumbnail":    u.UserProfile.Thumbnail,
		"created_at":   u.CreatedAt,
	}
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

func (user *User) RefreshUserToken(tokenId string, refreshTokenExp int64, originalRefreshToken string) helpers.JSON {
	now := time.Now().Unix()
	diff := refreshTokenExp - now

	log.Println("refreshing..")
	refreshToken := originalRefreshToken

	if diff < 60*60*24*15 {
		log.Println("refreshing refreshToken")
		accessSubject := "access_token"
		accessPayload := helpers.JSON{
			"user_id": user.ID,
		}

		accessToken, _ := helpers.GenerateAccessToken(accessPayload, accessSubject)

		refreshSubject := "refresh_token"
		refreshPayload := helpers.JSON{
			"user_id":  user.ID,
			"token_id": tokenId,
		}

		refreshToken, _ = helpers.GenerateRefreshToken(refreshPayload, refreshSubject)

		return helpers.JSON{
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		}
	}

	return nil
}
