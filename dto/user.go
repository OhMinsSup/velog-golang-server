package dto

import "time"

type UserRawQueryResult struct {
	ID          string    `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	ShortBio    string    `json:"short_bio"`
	Thumbnail   string    `json:"thumbnail"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
