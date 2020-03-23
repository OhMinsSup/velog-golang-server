package models

import "time"

type UserProfile struct {
	ID          string `gorm:"primary_key;uuid"`
	DisplayName string
	ShortBio    string
	Thumbnail   string
	UserID      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `sql:"index"`
}
