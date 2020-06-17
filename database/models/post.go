package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"time"
)

type Post struct {
	ID         string     `gorm:"primary_key;uuid"json:"id"`
	Title      string     `json:"title"`
	Body       string     `gorm:"type:text"json:"body"`
	Thumbnail  string     `json:"thumbnail"`
	IsMarkdown bool       `json:"is_markdown"`
	IsTemp     bool       `json:"is_temp"`
	IsPrivate  bool       `gorm:"default:true"json:"is_private"`
	UrlSlug    string     `sql:"index"json:"url_slug"`
	Likes      int        `gorm:"default:0"json:"likes"`
	Views      int        `gorm:"default:0"json:"views"`
	User       User       `gorm:"foreignkey:UserID"json:"user"`
	UserID     string     `json:"user_id"`
	CreatedAt  time.Time  `gorm:"type:time"json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:time"json:"updated_at"`
	DeletedAt  *time.Time `gorm:"type:time"sql:"index";json:"deleted_at"`
	Tags       []Tag      `gorm:"many2many:posts_tags;association_jointable_foreignkey:tag_id;jointable_foreignkey:post_id;"json:"tags"`
}

// Serialize serializes post data
func (p Post) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":         p.ID,
		"title":      p.Title,
		"body":       p.Body,
		"thumbnail":  p.Thumbnail,
		"url_slug":   p.UrlSlug,
		"likes":      p.Likes,
		"views":      p.Views,
		"is_temp":    p.IsTemp,
		"is_private": p.IsPrivate,
		"created_at": p.CreatedAt,
	}
}
