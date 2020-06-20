package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"time"
)

type UserMeta struct {
	ID                string     `gorm:"primary_key;uuid"json:"id"`
	EmailNotification bool       `gorm:"default:false"json:"email_notification"`
	EmailPromotion    bool       `gorm:"default:false"json:"email_promotion"`
	UserID            string     `json:"user_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}

func (u UserMeta) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":                 u.ID,
		"email_notification": u.EmailNotification,
		"email_promotion":    u.EmailPromotion,
	}
}
