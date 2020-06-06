package models

import "time"

type UserProfile struct {
	ID          string     `gorm:"primary_key;uuid"json:"id"`
	DisplayName string     `json:"display_name"`
	ShortBio    string     `json:"short_bio"`
	Thumbnail   string     `json:"thumbnail"`
	UserID      string     `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index"json:"deleted_at"`
}
