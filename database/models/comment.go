package models

import "time"

type Comment struct {
	ID         string     `gorm:"primary_key;uuid"json:"id"`
	Text       string     `json:"text"`
	Likes      int64      `gorm:"default:0"json:"likes"`
	ReplyTo    string     `json:"reply_to"`
	HasReplies bool       `gorm:"default:0"json:"has_replies"`
	Deleted    bool       `gorm:"default:0"json:"deleted"`
	CreatedAt  time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"gorm:"type:time"json:"deleted_at"`
}
