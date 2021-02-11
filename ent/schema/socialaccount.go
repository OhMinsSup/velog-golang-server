package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// SocialAccount holds the schema definition for the SocialAccount entity.
type SocialAccount struct {
	ent.Schema
}

// Fields of the SocialAccount.
func (SocialAccount) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("social_id").MaxLen(255),
		field.String("access_token").MaxLen(255),
		field.String("provider").MaxLen(255),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the SocialAccount.
func (SocialAccount) Edges() []ent.Edge {
	return []ent.Edge {
		edge.From("user", User.Type).
			Ref("social_account").
			Unique().
			Required(),
	}
}
