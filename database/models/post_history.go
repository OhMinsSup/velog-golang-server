package models

import "time"

type PostHistory struct {
	ID         string     `gorm:"primary_key;uuid"json:"id"`
	Title      string     `json:"title"`
	Body       string     `gorm:"type:text"json:"body"`
	IsMarkdown bool       `json:"is_markdown"`
	PostId     string     `sql:"index"json:"post_id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"json:"deleted_at"`
}
