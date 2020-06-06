package models

import (
	"time"
)

type PostsTags struct {
	ID        string     `gorm:"primary_key;uuid",json:"id"`
	TagId     string     `sql:"index"json:"tag_id"`
	PostId    string     `sql:"index"json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
}
