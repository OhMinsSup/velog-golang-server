package dto

import "github.com/OhMinsSup/story-server/libs"

// WritePostDTO - 포스트 작성 body 데이터
type WritePostDTO struct {
	Title      string    `json:"title" binding:"required"`
	Body       string    `json:"body" binding:"required"`
	IsMarkdown bool      `json:"is_markdown"`
	IsTemp     bool      `json:"is_temp"`
	IsPrivate  bool      `json:"is_private"`
	UrlSlug    string    `json:"url_slug"`
	Thumbnail  string    `json:"thumbnail"`
	Meta       libs.JSON `json:"meta"`
	Tags       []string  `json:"tag"`
}
