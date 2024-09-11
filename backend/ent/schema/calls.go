package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// CALLS holds the schema definition for the CALLS entity.
type CALLS struct {
	ent.Schema
}

// Fields of the CALLS.
func (CALLS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("session_id"),
		field.Time("call_start").
			Default(time.Now),
		field.Time("call_end").
			Default(time.Now),
		field.Int("rating"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the CALLS.
func (CALLS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("made", MATCHINGS.Type).
			Unique().
			Ref("makes"),
	}
}
