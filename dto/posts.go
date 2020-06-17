package dto

import (
	"github.com/lib/pq"
	"time"
)

type ListPostQuery struct {
	Cursor   string `json:"cursor"`
	Limit    int64  `json:"limit"`
	Username string `json:"username"`
}

type WritePostBody struct {
	Title      string   `json:"title"binding:"required"`
	Body       string   `json:"body"`
	Thumbnail  string   `json:"thumbnail"`
	IsMarkdown bool     `json:"is_markdown"`
	IsTemp     bool     `json:"is_temp"`
	IsPrivate  bool     `json:"is_private"`
	UrlSlug    string   `json:"url_slug"`
	Tag        []string `json:"tag"`
}

type PostRawQueryResult struct {
	ID         string         `json:"id"`
	Title      string         `json:"title"`
	Body       string         `json:"body"`
	Thumbnail  string         `json:"thumbnail"`
	IsMarkdown bool           `json:"is_markdown"`
	IsTemp     bool           `json:"is_temp"`
	IsPrivate  bool           `json:"is_private"`
	UrlSlug    string         `json:"url_slug"`
	Likes      int            `json:"likes"`
	Views      int            `json:"views"`
	UserID     string         `json:"user_id"`
	Tag        pq.StringArray `json:"tag"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  *time.Time     `json:"deleted_at"`
}

type PostsRawQueryResult struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Body          string         `json:"body"`
	Thumbnail     string         `json:"thumbnail"`
	IsMarkdown    bool           `json:"is_markdown"`
	IsTemp        bool           `json:"is_temp"`
	IsPrivate     bool           `json:"is_private"`
	UrlSlug       string         `json:"url_slug"`
	Likes         int            `json:"likes"`
	Views         int            `json:"views"`
	UserID        string         `json:"user_id"`
	Email         string         `json:"email"`
	Username      string         `json:"username"`
	DisplayName   string         `json:"display_name"`
	ShortBio      string         `json:"short_bio"`
	UserThumbnail string         `json:"user_thumbnail"`
	Tag           pq.StringArray `json:"tag"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     *time.Time     `json:"deleted_at"`
}
