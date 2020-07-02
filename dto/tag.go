package dto

import "time"

type SortingTagListQuery struct {
	Cursor string `json:"cursor"`
	Limit  int64  `json:"limit"`
	Sort   string `json:"sort"`
}

type TrendingTags struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	PostsCount int64     `json:"posts_count"`
	CreatedAt  time.Time `json:"created_at"`
}
