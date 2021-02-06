package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// UserProfile holds the schema definition for the UserProfile entity.
type UserProfile struct {
	ent.Schema
}

// Fields of the UserProfile.
func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("display_name").MaxLen(255),
		field.String("short_bio").MaxLen(255),
		field.Text("about").Optional().Nillable(),
		field.JSON("profile_links", []string{}).Optional(),
		field.String("thumbnail").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the UserProfile.
func (UserProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_profile").
			Unique().
			Required(),
	}
}
