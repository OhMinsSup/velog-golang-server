package models

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	ID        string     `gorm:"primary_key;uuid"json:"id"`
	Name      string     `sql:"index"json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	//DeletedAt *time.Time `sql:"index"json:"deleted_at"`
}

func (t Tag) Serialize() helpers.JSON {
	return helpers.JSON{
		"id":   t.ID,
		"name": t.Name,
	}
}

func TagFindOnCreate(name string, db *gorm.DB) Tag {
	var tagModel Tag
	if err := db.Where("name = ?", name).First(&tagModel).Error; err == nil {
		return tagModel
	}

	tagModel = Tag{
		Name: name,
	}

	db.NewRecord(tagModel)
	db.Create(&tagModel)
	return tagModel
}
