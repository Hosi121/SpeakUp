package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SESSIONS holds the schema definition for the SESSIONS entity.
type EVENTS struct {
	ent.Schema
}

// Fields of the SESSIONS.
func (EVENTS) Fields() []ent.Field {
	return []ent.Field{
		field.Time("event_start"),
		field.Time("event_end"),
		field.Int("theme_id"),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the SESSIONS.
func (EVENTS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participated", EVENT_RECORDS.Type),
		edge.To("uses", AI_THEMES.Type).
			Unique(),
	}
}
