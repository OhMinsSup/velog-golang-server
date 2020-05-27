package dto

type WritePostBody struct {
	Title      string             `json:"title", binding:"required"`
	Body       string             `json:"body"`
	Thumbnail  string             `json:"thumbnail"`
	IsMarkdown bool               `json:"is_markdown"`
	IsTemp     bool               `json:"is_temp"`
	IsPrivate  bool               `json:"is_private"`
	UrlSlug    string             `json:"url_slug"`
	Tag        []string           `json:"tag"`
}
