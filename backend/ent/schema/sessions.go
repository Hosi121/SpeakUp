package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SESSIONS holds the schema definition for the SESSIONS entity.
type SESSIONS struct {
	ent.Schema
}

// Fields of the SESSIONS.
func (SESSIONS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("session_id").
			Unique().
			Immutable(),
		field.Time("session_start"),
		field.Time("session_end"),
		field.Int("theme_id"),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the SESSIONS.
func (SESSIONS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("matchings", MATCHINGS.Type),
		edge.To("aithemes", AITHEMES.Type).
			Unique(),
		edge.From("included", CALLS.Type).
			Ref("includes"),
		edge.From("created", USERS.Type).Ref("creates"),
	}
}
