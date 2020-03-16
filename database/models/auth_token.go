package models

import "time"

type AuthToken struct {
	ID        string `gorm:"primary_key;uuid"`
	disabled  bool   `gorm:"default:false"`
	User      User   `gorm:"foreignkey:UserID"`
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
