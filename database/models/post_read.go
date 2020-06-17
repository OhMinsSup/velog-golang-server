package models

import "time"

type PostRead struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	IpHash    string     `sql:"index";json:"ip_hash"`
	UserId    string     `sql:"index"json:"user_id"`
	PostId    string     `sql:"index"json:"post_id"`
	CreatedAt time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:time"sql:"index";json:"deleted_at"`
}
