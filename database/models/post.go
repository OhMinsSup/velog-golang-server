package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type MetaPayload map[string]interface{}

func (meta MetaPayload) Value() (driver.Value, error) {
	j, err := json.Marshal(meta)
	return j, err
}

func (meta *MetaPayload) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	if err := json.Unmarshal(source, &i); err != nil {
		return err
	}

	*meta, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}

type Post struct {
	ID         string `gorm:"primary_key;uuid"`
	Title      string
	Body       string `gorm:"type:text"`
	Thumbnail  string
	IsMarkdown bool
	IsTemp     bool
	IsPrivate  bool        `gorm:"default:true"`
	UrlSlug    string      `sql:"index"`
	Likes      int         `gorm:"default:0"`
	Views      int         `gorm:"default:0"`
	Meta       MetaPayload `gorm:"type:json;not null;default '{}'"`
	User       User        `gorm:"foreignkey:UserID"`
	UserID     string
	ReleasedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `sql:"index"`
	Tags       []Tag      `gorm:"many2many:PostsTags;association_jointable_foreignkey:tag_id;jointable_foreignkey:post_id;"`
}
