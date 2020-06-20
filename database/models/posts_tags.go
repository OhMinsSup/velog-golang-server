package models

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers/fx"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"time"
)

type PostsTags struct {
	ID        string     `gorm:"primary_key;uuid",json:"id"`
	TagId     string     `sql:"index"json:"tag_id"`
	PostId    string     `sql:"index"json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	//DeletedAt *time.Time `sql:"index"json:"deleted_at"`
}

func SyncPostTags(body dto.WritePostBody, post Post, db *gorm.DB) {
	var tagIds []string
	for iter := 0; iter < len(body.Tag); iter++ {
		currentTag := body.Tag[iter]
		tag := TagFindOnCreate(currentTag, db)
		tagIds = append(tagIds, tag.ID)
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

	var prevPostTag struct{
		TagIds pq.StringArray `json:"tag_ids"`
	}
	// current posts tags info
	if err := db.Raw(`
		SELECT DISTINCT array_agg(pt.tag_id) AS tag_ids FROM posts p
		INNER JOIN posts_tags pt ON pt.post_id = p.id
		WHERE pt.post_id = ?
		GROUP BY p.id, pt.post_id`, post.ID).Find(&prevPostTag).Error; err != nil {
		panic(err)
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
			if err := db.Raw("DELETE FROM posts_tags pt WHERE pt.tag_id = ? AND pt.post_id = ?", missingTagId, post.ID).Error; err != nil {
				panic(err)
			}
		}
	}

	// adding tags
	if len(adding) > 0 {
		for _, addingTagId := range adding {
			postsTags := PostsTags{
				PostId: post.ID,
				TagId:  addingTagId,
			}

			db.NewRecord(postsTags)
			db.Create(&postsTags)
		}
	}
}
