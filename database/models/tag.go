package models

import "time"

type Tag struct {
	ID        string `gorm:"primary_key;uuid"`
	Name      string `sql:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
