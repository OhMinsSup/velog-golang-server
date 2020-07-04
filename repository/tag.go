package repository

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/models"
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

func (t *TagRepository) GetPostsCount(tagId string) (int64, error) {
	if tagId == "" {
		return 0, nil
	}

	type Count struct {
		PostsCount int64 `json:"posts_count"`
	}
	var count []Count
	if err := t.db.Raw(`
		SELECT posts_count FROM (
		  SELECT count(post_id) AS posts_count, posts_tags.tag_id AS tag_id FROM posts_tags
		  INNER JOIN posts ON posts.id = post_id
		  WHERE posts.is_private = FALSE
		  AND posts.is_temp = FALSE
		  GROUP BY tag_id
		) AS q1
		WHERE tag_id = ?`, tagId).Scan(&count).Error; err != nil {
		return 0, err
	}

	if len(count) == 0 {
		return 0, nil
	}

	return count[0].PostsCount, nil
}

func (t *TagRepository) GetTagList(cursor string, limit int64) ([]dto.Tags, error) {
	if cursor == "" {
		var tags []dto.Tags
		if err := t.db.Raw(`
		SELECT tags.id, tags.name, tags.created_at, posts_count FROM (
			SELECT count(post_id) AS posts_count, posts_tags.tag_id AS tag_id FROM posts_tags 
			INNER JOIN posts ON posts.id = post_id
			WHERE posts.is_private = FALSE
			AND posts.is_temp = FALSE
			GROUP BY tag_id
		) AS q1
		INNER JOIN tags ON q1.tag_id = tags.id
		ORDER BY tags.name
		LIMIT ?`, limit).Scan(&tags).Error; err != nil {
			return nil, err
		}
		return tags, nil
	}

	var cursorTag models.Tag
	if err := t.db.Where("id = ?", cursor).First(&cursorTag).Error; err != nil {
		return nil, err
	}

	var tags []dto.Tags
	if err := t.db.Raw(`
		SELECT tags.id, tags.name, tags.created_at, posts_count FROM (
			SELECT count(post_id) AS posts_count, posts_tags.tag_id AS tag_id FROM posts_tags
			INNER JOIN posts ON posts.id = post_id
			WHERE posts.is_private = FALSE
			AND posts.is_temp = FALSE
			GROUP BY tag_id
		) AS q1
		INNER JOIN tags ON q1.tag_id = tags.id
		where tags.name > ?
		ORDER BY tags.name
		LIMIT ?`, cursorTag.Name, limit).Scan(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t *TagRepository) TrendingTagList(cursor string, limit int64) ([]dto.Tags, error) {
	if cursor == "" {
		var tags []dto.Tags
		if err := t.db.Raw(`
		SELECT tags.id, tags.name, tags.created_at, posts_count FROM (
			SELECT count(post_id) AS posts_count, posts_tags.tag_id AS tag_id FROM posts_tags
			INNER JOIN posts ON posts.id = post_id
			WHERE posts.is_private = FALSE
			AND posts.is_temp = FALSE
			GROUP BY tag_id
		) AS q1
		INNER JOIN tags ON q1.tag_id = tags.id
		ORDER BY posts_count desc, tags.id
		LIMIT ?`, limit).Scan(&tags).Error; err != nil {
			return nil, err
		}
		return tags, nil
	}

	cursorPostsCount, err := t.GetPostsCount(cursor)
	if err != nil {
		return nil, err
	}

	var tags []dto.Tags
	if err := t.db.Raw(`
		SELECT tags.id, tags.name, tags.created_at, posts_count FROM (
			SELECT count(post_id) AS posts_count, posts_tags.tag_id AS tag_id FROM posts_tags
			INNER JOIN posts ON posts.id = post_id
			WHERE posts.is_private = FALSE
			AND posts.is_temp = FALSE
			GROUP BY tag_id
		) AS q1
		INNER JOIN tags ON q1.tag_id = tags.id
        WHERE posts_count <= ?
        AND id != ?
        AND NOT (id < ? AND posts_count = ?)
		ORDER BY posts_count desc, tags.id
		LIMIT ?`, cursorPostsCount, cursor, cursor, cursorPostsCount, limit).Scan(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
