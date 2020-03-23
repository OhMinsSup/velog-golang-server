package models

import "time"

type VelogConfig struct {
	ID        string `gorm:"primary_key;uuid"`
	title     string
	logoImage string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
