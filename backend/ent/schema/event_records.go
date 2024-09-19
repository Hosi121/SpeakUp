package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AITHEMES holds the schema definition for the AITHEMES entity.
type EVENT_RECORDS struct {
	ent.Schema
}

// Fields of the AITHEMES.
func (EVENT_RECORDS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("event_id"),
		field.Int("participates_bit").
			Max(7).
			Min(0).
			Default(0),
		field.String("records").
			Default(""),
	}
}

// Edges of the AITHEMES.
func (EVENT_RECORDS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("made", USERS.Type).
			Ref("makes").
			Unique(),
		edge.From("participates", EVENTS.Type).
			Ref("participated").
			Unique(),
		edge.To("has", SESSIONS.Type),
	}
}
