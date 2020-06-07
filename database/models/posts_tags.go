package models

import (
	"github.com/OhMinsSup/story-server/dto"
	"github.com/OhMinsSup/story-server/helpers/fx"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

type PostsTags struct {
	ID        string     `gorm:"primary_key;uuid",json:"id"`
	TagId     string     `sql:"index"json:"tag_id"`
	PostId    string     `sql:"index"json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
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

	type PrevPostTags struct {
		TagId string `json:"tag_id"`
	}

	var prevPostTags []PrevPostTags
	// 현재 포스트에 등록된 태그 정보
	db.Raw("SELECT DISTINCT pt.tag_id FROM posts p INNER JOIN posts_tags pt ON pt.post_id = p.id WHERE pt.post_id = ?", post.ID).Find(&prevPostTags)

	// get deleted posts_tags Item
	var missing []string
	for _, pt := range prevPostTags {
		if id, prefix := fx.ContainSelector(tagIds, pt.TagId); prefix {
			log.Println("missing", id)
			missing = append(missing, id)
		}
	}

	// get add posts_tags Item
	var adding []string
	for _, t := range tagIds {
		if len(prevPostTags) > 0 {
			for _, pt := range prevPostTags {
				if !strings.Contains(t, pt.TagId) {
					adding = append(adding, t)
				}
			}
		} else {
			adding = append(adding, t)
		}
	}

	// remove tags
	if len(missing) > 0 {
		for _, missingTagId := range missing {
			db.Raw("DELETE FROM posts_tags pt WHERE pt.tag_id = ? AND pt.post_id = ?", missingTagId, post.ID)
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
