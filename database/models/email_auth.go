package models

import "time"

type EmailAuth struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Code      string     `sql:"index"json:"code"`
	Email     string     `json:"email"`
	Logged    bool       `gorm:"default:false"json:"logged"`
	CreatedAt time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:time"sql:"index"json:"deleted_at"`
}
