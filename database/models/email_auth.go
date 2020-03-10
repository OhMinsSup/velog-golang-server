package models

import "time"

type EmailAuth struct {
	ID        string `gorm:"primary_key"`
	Code      string `sql:"index"`
	Email     string
	Logged    bool `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
