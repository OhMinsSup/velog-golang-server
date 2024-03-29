package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// UserMeta holds the schema definition for the UserMeta entity.
type UserMeta struct {
	ent.Schema
}

// Fields of the UserMeta.
func (UserMeta) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bool("email_notification").Default(false),
		field.Bool("email_promotions").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserMeta.
func (UserMeta) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_meta").
			Unique().
			Required(),
	}
}
