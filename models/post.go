package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"time"
)

type Post struct {
	ID          string        `gorm:"primary_key;uuid"json:"id"`
	Title       string        `json:"title"`
	Body        string        `gorm:"type:text"json:"body"`
	Thumbnail   string        `json:"thumbnail"`
	IsMarkdown  bool          `json:"is_markdown"`
	IsTemp      bool          `json:"is_temp"`
	IsPrivate   bool          `json:"is_private"`
	Likes       int64         `gorm:"default:0"json:"likes"`
	Views       int64         `gorm:"default:0"json:"views"`
	User        User          `gorm:"foreignkey:UserID"json:"user"`
	UserID      string        `json:"user_id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Tags        []Tag         `gorm:"many2many:posts_tags;association_autocreate:false"json:"tags"`
	PostScore   []PostScore   `gorm:"polymorphic:Owner;"json:"post_score"`
	PostRead    []PostRead    `gorm:"polymorphic:Owner;"json:"post_read"`
	PostHistory []PostHistory `gorm:"polymorphic:Owner;"json:"post_history"`
	PostLike    []PostLike    `gorm:"polymorphic:Owner;"json:"post_like"`
	PostComment []Comment     `gorm:"polymorphic:Owner;"json:"post_comment"`
}

func (p Post) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":          p.ID,
		"user_id":     p.UserID,
		"title":       p.Title,
		"body":        p.Body,
		"thumbnail":   p.Thumbnail,
		"likes":       p.Likes,
		"views":       p.Views,
		"is_markdown": p.IsMarkdown,
		"is_temp":     p.IsTemp,
		"is_private":  p.IsPrivate,
		"created_at":  p.CreatedAt,
		"updated_at":  p.UpdatedAt,
	}
}

type Tag struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	Name      string    `sql:"index"json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Tag) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":   t.ID,
		"name": t.Name,
	}
}

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

type PostRead struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	IpHash    string    `sql:"index";json:"ip_hash"`
	UserId    string    `sql:"index"json:"user_id"`
	PostId    string    `sql:"index"json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostLike struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	UserId    string    `sql:"index"json:"user_id"`
	PostId    string    `sql:"index"json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostScore struct {
	ID        string    `gorm:"primary_key;uuid"json:"id"`
	Type      string    `json:"type"`
	Score     float64   `sql:"type:decimal(10,2);"json:"score"`
	UserId    string    `sql:"index"json:"user_id"`
	PostId    string    `sql:"index"json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostHistory struct {
	ID         string    `gorm:"primary_key;uuid"json:"id"`
	Title      string    `json:"title"`
	Body       string    `gorm:"type:text"json:"body"`
	IsMarkdown bool      `json:"is_markdown"`
	PostId     string    `sql:"index"json:"post_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
