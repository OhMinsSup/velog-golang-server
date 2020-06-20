package models

import "time"

type AuthToken struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Disabled  bool       `gorm:"default:false"json:"disabled"`
	User      User       `gorm:"foreignkey:UserID"json:"user"`
	UserID    string     `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
