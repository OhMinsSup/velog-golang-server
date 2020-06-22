package models

import "time"

type Comment struct {
	ID         string    `gorm:"primary_key;uuid"json:"id"`
	Text       string    `json:"text"`
	Likes      int64     `gorm:"default:0"json:"likes"`
	UserId     string    `sql:"index"json:"user_id"`
	PostId     string    `sql:"index"json:"post_id"`
	ReplyTo    string    `json:"reply_to"`
	HasReplies bool      `gorm:"default:false"json:"has_replies"`
	Deleted    bool      `gorm:"default:false"json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
