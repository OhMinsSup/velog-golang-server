package models

import "time"

type EmailAuth struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Code      string     `sql:"index"json:"code"`
	Email     string     `json:"email"`
	Logged    bool       `gorm:"default:false"json:"logged"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
