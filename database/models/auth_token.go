package models

import "time"

type AuthToken struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Disabled  bool       `gorm:"default:false"json:"disabled"`
	User      User       `gorm:"foreignkey:UserID"json:"user"`
	UserID    string     `json:"user_id"`
	CreatedAt time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt *time.Time `sql:"index"gorm:"type:time"json:"deleted_at"`
}
