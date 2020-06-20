package models

import "time"

type PostScore struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Type      string     `json:"type"`
	Score     float64    `sql:"type:decimal(10,2);"json:"score"`
	UserId    string     `sql:"index"json:"user_id"`
	PostId    string     `sql:"index"json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	//DeletedAt *time.Time `sql:"index"json:"deleted_at"`
}
