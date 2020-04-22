package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	ID        string `gorm:"primary_key;uuid"`
	Name      string `sql:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
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
