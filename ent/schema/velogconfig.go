package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// VelogConfig holds the schema definition for the VelogConfig entity.
type VelogConfig struct {
	ent.Schema
}

// Fields of the VelogConfig.
func (VelogConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("title").Nillable().Optional(),
		field.String("logo_title").Nillable().Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the VelogConfig.
func (VelogConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("velog_config").
			Unique().
			Required(),
	}
}
