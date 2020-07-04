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
	AuthTokens  []AuthToken `gorm:"foreignkey:UserID"json:"auth_tokens"`
	UserProfile UserProfile `gorm:"foreignkey:UserID"json:"user_profile"`
	UserMeta    UserMeta    `gorm:"foreignkey:UserID"json:"user_meta"`
	VelogConfig VelogConfig `gorm:"foreignkey:UserID"json:"velog_config"`
	PostScore   []PostScore `gorm:"polymorphic:Owner;"json:"post_score"`
	PostRead    []PostRead  `gorm:"polymorphic:Owner;"json:"post_read"`
	PostLike    []PostLike  `gorm:"polymorphic:Owner;"json:"post_like"`
	PostComment []Comment   `gorm:"polymorphic:Owner;"json:"post_comment"`
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

func (u *User) GenerateUserToken(db *gorm.DB) helpers.JSON {
	authToken := AuthToken{
		UserID: u.ID,
	}

	db.NewRecord(authToken)
	db.Create(&authToken)

	accessSubject := "access_token"
	accessPayload := helpers.JSON{
		"user_id": u.ID,
	}

	accessToken, _ := helpers.GenerateAccessToken(accessPayload, accessSubject)

	refreshSubject := "refresh_token"
	refreshPayload := helpers.JSON{
		"user_id":  u.ID,
		"token_id": authToken.ID,
	}

	refreshToken, _ := helpers.GenerateRefreshToken(refreshPayload, refreshSubject)

	return helpers.JSON{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
}

func (u *User) RefreshUserToken(tokenId string, refreshTokenExp int64, originalRefreshToken string) helpers.JSON {
	now := time.Now().Unix()
	diff := refreshTokenExp - now

	log.Println("refreshing..")
	refreshToken := originalRefreshToken

	if diff < 60*60*24*15 {
		log.Println("refreshing refreshToken")
		accessSubject := "access_token"
		accessPayload := helpers.JSON{
			"user_id": u.ID,
		}

		accessToken, _ := helpers.GenerateAccessToken(accessPayload, accessSubject)

		refreshSubject := "refresh_token"
		refreshPayload := helpers.JSON{
			"user_id":  u.ID,
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

type UserProfile struct {
	ID          string    `gorm:"primary_key;uuid"json:"id"`
	DisplayName string    `json:"display_name"`
	ShortBio    string    `json:"short_bio"`
	Thumbnail   string    `json:"thumbnail"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u UserProfile) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":           u.ID,
		"display_name": u.DisplayName,
		"short_bio":    u.ShortBio,
		"thumbnail":    u.Thumbnail,
	}
}

type UserMeta struct {
	ID                string    `gorm:"primary_key;uuid"json:"id"`
	EmailNotification bool      `gorm:"default:false"json:"email_notification"`
	EmailPromotion    bool      `gorm:"default:false"json:"email_promotion"`
	UserID            string    `json:"user_id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (u UserMeta) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":                 u.ID,
		"email_notification": u.EmailNotification,
		"email_promotion":    u.EmailPromotion,
	}
}

type EmailAuth struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	Code      string    `sql:"index"json:"code"`
	Email     string    `json:"email"`
	Logged    bool      `gorm:"default:false"json:"logged"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthToken struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	Disabled  bool      `gorm:"default:false"json:"disabled"`
	User      User      `gorm:"foreignkey:UserID"json:"user"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SocialAccount struct {
	ID          string    `gorm:"primary_key;uuid"json:"id"`
	SocialId    string    `sql:"index"json:"social_id"`
	AccessToken string    `json:"access_token"`
	Provider    string    `sql:"index"json:"provider"`
	User        User      `gorm:"foreignkey:UserID"json:"user"`
	UserID      string    `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type VelogConfig struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	Title     string    `json:"title"`
	LogoImage string    `json:"logo_image"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (v VelogConfig) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":         v.ID,
		"title":      v.Title,
		"logo_image": v.LogoImage,
	}
}
