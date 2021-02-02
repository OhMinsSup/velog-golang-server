package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/google/uuid"
	"time"
)

// EmailAuth holds the schema definition for the EmailAuth entity.
type EmailAuth struct {
	ent.Schema
}

// Fields of the EmailAuth.
func (EmailAuth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("code"),
		field.String("email"),
		field.Bool("logged").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Indexes of the EmailAuth.
func (EmailAuth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("code"),
	}
}

// Edges of the EmailAuth.
func (EmailAuth) Edges() []ent.Edge {
	return nil
}
