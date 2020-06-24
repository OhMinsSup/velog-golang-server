package models

import "time"

type PostLike struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	UserId    string     `sql:"index"json:"user_id"`
	PostId    string     `sql:"index"json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}