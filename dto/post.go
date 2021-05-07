package dto

import (
	"github.com/OhMinsSup/story-server/libs"
	"github.com/google/uuid"
	"time"
)

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
	Tags       []string  `json:"tags"`
}

// UpdatePostDTO - 포스트 수정 body 데이터
type UpdatePostDTO struct {
	PostID     string    `json:"post_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body" binding:"required"`
	IsMarkdown bool      `json:"is_markdown"`
	IsTemp     bool      `json:"is_temp"`
	IsPrivate  bool      `json:"is_private"`
	UrlSlug    string    `json:"url_slug"`
	Thumbnail  string    `json:"thumbnail"`
	Meta       libs.JSON `json:"meta"`
	Tags       []string  `json:"tags"`
}

// ListPostQueryString - 포스트 리스트의 쿼리스트링
type ListPostQueryString struct {
	Cursor   string `json:"cursor"`
	Limit    int64  `json:"limit"`
	Username string `json:"username"`
	TempOnly bool   `json:"temp_only"`
	Tag      string `json:"tag"`
}

// PostTagSchema - 포스트 일기 상태에서 태그
type PostTagSchema struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// PostUserProfileSchema - 포스트 읽기 상세에서 유저 프로필 데이터
type PostUserProfileSchema struct {
	Thumbnail string `json:"thumbnail"`
	ShortBio  string `json:"short_bio"`
}

// PostUserSchema - 포스트 읽기 상세에서 유저 데이터
type PostUserSchema struct {
	ID          uuid.UUID             `json:"id"`
	Username    string                `json:"username"`
	Email       string                `json:"email"`
	UserProfile PostUserProfileSchema `json:"user_profile"`
}

// ReadPostSchema - 포스트 읽기 상세 데이터
type ReadPostSchema struct {
	ID         uuid.UUID       `json:"id"`
	Title      string          `json:"title"`
	Body       string          `json:"body"`
	Thumbnail  string          `json:"thumbnail"`
	IsMarkdown bool            `json:"is_markdown"`
	IsTemp     bool            `json:"is_temp"`
	IsPrivate  bool            `json:"is_private"`
	UrlSlug    string          `json:"url_slug"`
	Likes      int64           `json:"likes"`
	Views      int64           `json:"views"`
	ReleasedAt time.Time       `json:"released_at"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	User       PostUserSchema  `json:"user"`
	Tags       []PostTagSchema `json:"tags"`
}
