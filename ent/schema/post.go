package schema

import (
	"github.com/OhMinsSup/story-server/libs"
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("fk_user_id", uuid.UUID{}),
		field.String("title").MaxLen(255),
		field.Text("body"),
		field.String("thumbnail").Nillable().Optional().MaxLen(255),
		field.Bool("is_temp"),
		field.Bool("is_markdown"),
		field.Bool("is_private").Default(true),
		field.String("url_slug").Unique().MaxLen(255),
		field.Int64("likes").Default(0),
		field.Int64("views").Default(0),
		field.JSON("meta", libs.JSON{}),
		field.Time("released_at").Default(time.Now),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("posts").Unique(),
	}
}

// Indexs of the Post.
func (Post) Indexs() []ent.Index {
	return []ent.Index{
		index.Fields("url_slug"),
	}
}
