package models

import "time"

type Comment struct {
	ID         string     `gorm:"primary_key;uuid"json:"id"`
	Text       string     `json:"text"`
	Likes      int64      `gorm:"default:0"json:"likes"`
	ReplyTo    string     `json:"reply_to"`
	HasReplies bool       `gorm:"default:0"json:"has_replies"`
	Deleted    bool       `gorm:"default:0"json:"deleted"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"json:"deleted_at"`
}
