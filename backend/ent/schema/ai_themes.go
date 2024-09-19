package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AITHEMES holds the schema definition for the AITHEMES entity.
type AI_THEMES struct {
	ent.Schema
}

// Fields of the AITHEMES.
func (AI_THEMES) Fields() []ent.Field {
	return []ent.Field{
		field.String("theme_text"),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the AITHEMES.
func (AI_THEMES) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("used", EVENTS.Type).
			Ref("uses"),
	}
}
