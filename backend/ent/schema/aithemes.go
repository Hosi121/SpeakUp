package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AITHEMES holds the schema definition for the AITHEMES entity.
type AITHEMES struct {
	ent.Schema
}

// Fields of the AITHEMES.
func (AITHEMES) Fields() []ent.Field {
	return []ent.Field{
		field.String("theme_text"),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the AITHEMES.
func (AITHEMES) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("used", SESSIONS.Type).
			Ref("uses"),
	}
}
