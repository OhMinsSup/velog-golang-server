package models

import "time"

type UserMeta struct {
	ID                string `gorm:"primary_key;uuid"`
	EmailNotification bool   `gorm:"default:false"`
	EmailPromotion    bool   `gorm:"default:false"`
	UserID            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         *time.Time `sql:"index"`
}
