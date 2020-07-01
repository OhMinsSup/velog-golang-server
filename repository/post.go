package repository

import (
	"github.com/OhMinsSup/story-server/database/models"
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/OhMinsSup/story-server/helpers/fx"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (p *PostRepository) GetPost(postId string) (helpers.JSON, error) {
	var post dto.PostRawQueryResult
	if err := p.db.Raw(`
		SELECT
		p.*,
		array_agg(t.name) AS tag FROM "posts" AS p
		LEFT OUTER JOIN "posts_tags" AS pt  ON pt.post_id = p.id
		LEFT OUTER JOIN "tags" AS t ON t.id = pt.tag_id
		WHERE p.id = ?
		GROUP BY p.id, pt.post_id`, postId).Scan(&post).Error; err != nil {
		return nil, err
	}

	var user dto.UserRawQueryResult
	if err := p.db.Raw(`
       SELECT
	   u.*,
	   up.display_name,
       up.short_bio,
	   up.thumbnail
	   FROM "users" AS u
	   LEFT OUTER JOIN "user_profiles" AS up ON up.user_id = u.id
	   WHERE u.id = ?`, post.UserID).Scan(&user).Error; err != nil {
		return nil, err
	}

	post.User = user

	return helpers.JSON{
		"post": post,
	}, nil
}

func (p *PostRepository) CreatePost(body dto.WritePostBody, userId string) (string, error) {
	newPost := models.Post{
		Title:      body.Title,
		Body:       body.Body,
		Thumbnail:  body.Thumbnail,
		IsTemp:     body.IsTemp,
		IsMarkdown: body.IsMarkdown,
		IsPrivate:  body.IsPrivate,
		UserID:     userId,
	}

	tx := p.db.Begin()
	if err := tx.Create(&newPost).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	if err := p.SyncPostTags(body.Tag, newPost.ID, newPost); err != nil {
		tx.Rollback()
		return "", err
	}

	return newPost.ID, tx.Commit().Error
}

func (p *PostRepository) UpdatePost(body dto.WritePostBody, userId, postId string) (string, error) {
	tx := p.db.Begin()

	var currentPost models.Post
	if err := tx.Where("id = ?", postId).First(&currentPost).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	if currentPost.UserID != userId {
		return "", helpers.ErrorPermission
	}

	if err := tx.Model(&currentPost).Updates(models.Post{
		Title:      body.Title,
		Body:       body.Body,
		Thumbnail:  body.Thumbnail,
		IsMarkdown: body.IsMarkdown,
		IsPrivate:  body.IsPrivate,
		UserID:     userId,
	}).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	if len(body.Tag) > 0 {
		if err := p.SyncPostTags(body.Tag, postId, currentPost); err != nil {
			tx.Rollback()
			return "", err
		}
	}

	return postId, tx.Commit().Error
}

func (p *PostRepository) DeletePost(userId, postId string) (bool, error) {
	var currentPost models.Post
	if err := p.db.Where("id = ?", postId).Preload("Tags").First(&currentPost).Error; err != nil {
		return false, err
	}

	if currentPost.UserID != userId {
		return false, helpers.ErrorPermission
	}

	tx := p.db.Begin()

	if err := tx.Model(&currentPost).Association("Tags").Delete(&currentPost.Tags).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	if err := tx.Delete(&currentPost).Error; err != nil {
		tx.Rollback()
		return false, err
	}

	return true, tx.Commit().Error
}

func (p *PostRepository) View(body dto.PostViewParams, userId string) error {
	var currentRead models.PostRead
	if err := p.db.Where(`ip_hash = ? AND post_id = ?`, body.Ip, body.PostId).First(&currentRead).Error; err == nil {
		if currentRead == (models.PostRead{}) {
			return nil
		}
	}

	tx := p.db.Begin()

	postRead := models.PostRead{
		PostId: body.PostId,
		UserId: userId,
		IpHash: body.Ip,
	}

	if err := tx.Create(&postRead).Error; err != nil {
		tx.Rollback()
		return err
	}

	var updatePost models.Post
	if err := tx.Where("id = ? AND created_at > (NOW() - INTERVAL '24 HOURS')", body.PostId).First(&updatePost).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&updatePost).Update(map[string]interface{}{"views": updatePost.Views + 1}).Error; err != nil {
		tx.Rollback()
		return err
	}

	newPostScore := models.PostScore{
		Type:   "READ",
		PostId: body.PostId,
		UserId: userId,
		Score:  1.0,
	}

	if err := tx.Create(&newPostScore).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (p *PostRepository) SyncPostTags(tags []string, postId string, txPost models.Post) error {
	tagRepository := NewTagRepository(p.db)

	var tagIds []string
	for _, tag := range tags {
		tagId, err := tagRepository.FindTagAndCreate(tag)
		if err != nil {
			return err
		}

		tagIds = append(tagIds, tagId)
	}

	// 중복을 제거한 배열을 얻는다.
	var uniqueTagIds []string
	filterTagIds := make(map[string]bool)
	for _, value := range tagIds {
		if _, tagId := filterTagIds[value]; !tagId {
			filterTagIds[value] = true
			uniqueTagIds = append(uniqueTagIds, value)
		}
	}

	var prevPostTag struct {
		TagIds pq.StringArray `json:"tag_ids"`
	}

	// current posts tags info
	if err := p.db.Raw(`
		SELECT DISTINCT array_agg(pt.tag_id) AS tag_ids FROM posts p
		INNER JOIN posts_tags pt ON pt.post_id = p.id
		WHERE pt.post_id = ?
		GROUP BY p.id, pt.post_id`, postId).Find(&prevPostTag).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
		return err
	}

	// get deleted posts_tags Item
	var missing []string
	for i, prevTagId := range prevPostTag.TagIds {
		if _, prefix := fx.ContainSelector(tagIds, prevTagId); !prefix {
			missing = append(missing, prevPostTag.TagIds[i])
		}
	}

	// get add posts_tags Item
	var adding []string
	for i, tagId := range tagIds {
		if _, prefix := fx.ContainSelector(prevPostTag.TagIds, tagId); !prefix {
			adding = append(adding, tagIds[i])
		}
	}

	// remove tags
	if len(missing) > 0 {
		for _, missingTagId := range missing {
			var prevTag models.Tag
			if err := p.db.Where("id = ?", missingTagId).First(&prevTag).Error; err != nil {
				return err
			}

			if err := p.db.Model(&txPost).Association("Tags").Delete(prevTag).Error; err != nil {
				return err
			}
		}
	}

	// adding tags
	if len(adding) > 0 {
		for _, addingTagId := range adding {
			var newTag models.Tag
			if err := p.db.Where("id = ?", addingTagId).First(&newTag).Error; err != nil {
				return err
			}

			if err := p.db.Model(&txPost).Association("Tags").Append(newTag).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
