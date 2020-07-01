package repository

import (
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/jinzhu/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (t *TagRepository) FindTagAndCreate(name string) (string, error) {
	var tag models.Tag
	if err := t.db.FirstOrCreate(&tag, models.Tag{Name: name}).Error; err != nil {
		return "", err
	}
	return tag.ID, nil
}
