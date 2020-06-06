package models

import "time"

type UserMeta struct {
	ID                string     `gorm:"primary_key;uuid"json:"id"`
	EmailNotification bool       `gorm:"default:false"json:"email_notification"`
	EmailPromotion    bool       `gorm:"default:false"json:"email_promotion"`
	UserID            string     `json:"user_id"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	DeletedAt         *time.Time `sql:"index"json:"deleted_at"`
}
