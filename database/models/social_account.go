package models

import "time"

type SocialAccount struct {
	ID          string `gorm:"primary_key;uuid"`
	SocialId    string `sql:"index"`
	AccessToken string
	Provider    string `sql:"index"`
	User        User   `gorm:"foreignkey:UserID"`
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
