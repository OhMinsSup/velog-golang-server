package models

import (
	"time"
)

type Post struct {
	ID         string `gorm:"primary_key;uuid"`
	Title      string
	Body       string `gorm:"type:text"`
	Thumbnail  string
	IsMarkdown bool
	IsTemp     bool
	IsPrivate  bool        `gorm:"default:true"`
	UrlSlug    string      `sql:"index"`
	Likes      int         `gorm:"default:0"`
	Views      int         `gorm:"default:0"`
	User       User        `gorm:"foreignkey:UserID"`
	UserID     string
	ReleasedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	Tags       []Tag      `gorm:"many2many:PostsTags;association_jointable_foreignkey:tag_id;jointable_foreignkey:post_id;"`
}
