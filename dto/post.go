package dto

import "github.com/OhMinsSup/story-server/database/models"

type WritePostBody struct {
	Title      string             `json:"title", binding:"required"`
	Body       string             `json:"body"`
	Thumbnail  string             `json:"thumbnail"`
	IsMarkdown bool               `json:"is_markdown"`
	IsTemp     bool               `json:"is_temp"`
	IsPrivate  bool               `json:"is_private"`
	UrlSlug    string             `json:"url_slug"`
	Meta       models.MetaPayload `json:"meta"`
	Tag        []string           `json:"tag"`
}
