package models

import "time"

type VelogConfig struct {
	ID        string `gorm:"primary_key;uuid", json:"id"`
	title     string `json:"title"`
	logoImage string `json:"logo_image"`
	UserID    string `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `sql:"index", json:"deleted_at"`
}
