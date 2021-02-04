package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("username").Unique().MaxLen(255),
		field.String("email").Unique().Nillable().Optional().MaxLen(255),
		field.Bool("is_certified").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username", "email").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_profile", UserProfile.Type).
			StorageKey(edge.Column("fk_user_id")).Unique(),
		edge.To("velog_config", VelogConfig.Type).
			StorageKey(edge.Column("fk_user_id")).Unique(),
		edge.To("user_meta", UserMeta.Type).
			StorageKey(edge.Column("fk_user_id")).Unique(),
		edge.To("auth_token", AuthToken.Type).
			StorageKey(edge.Column("fk_user_id")),
	}
}
