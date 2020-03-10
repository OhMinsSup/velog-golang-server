package models

import "time"

type AuthToken struct {
	ID        string `gorm:"primary_key"`
	disabled  bool   `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
