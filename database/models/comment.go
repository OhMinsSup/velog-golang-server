package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"time"
)

type Comment struct {
	ID         string    `gorm:"primary_key;uuid"json:"id"`
	Text       string    `json:"text"`
	Likes      int64     `gorm:"default:0"json:"likes"`
	UserId     string    `sql:"index"json:"user_id"`
	PostId     string    `sql:"index"json:"post_id"`
	Level      int64     `gorm:"default:0"json:"level"`
	ReplyTo    string    `json:"reply_to"`
	HasReplies bool      `gorm:"default:false"json:"has_replies"`
	Deleted    bool      `gorm:"default:false"json:"deleted"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (c Comment) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":          c.ID,
		"user_id":     c.UserId,
		"post_id":     c.PostId,
		"text":        c.Text,
		"level":       c.Level,
		"reply_to":    c.ReplyTo,
		"has_replies": c.HasReplies,
		"deleted":     c.Deleted,
		"likes":       c.Likes,
		"created_at":  c.CreatedAt,
		"updated_at":  c.UpdatedAt,
	}
}
