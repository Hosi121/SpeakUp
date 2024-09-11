package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MATCHINGS holds the schema definition for the MATCHINGS entity.
type MATCHINGS struct {
	ent.Schema
}

// Fields of the MATCHINGS.
func (MATCHINGS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("matching_id").
			Unique().
			Immutable(),
		field.Int("user_id"),
		field.Int("matched_user_id"),
		field.Int("session_id"),
		field.Time("matched_at").
			Default(time.Now()),
		field.Enum("status").
			Values("MATCHED", "PROCCESSING", "FINISHED"),
	}
}

// Edges of the MATCHINGS.
func (MATCHINGS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", USERS.Type).
			Ref("participates").Unique(),
		edge.From("had", SESSIONS.Type).
			Ref("has").Unique(),
	}
}
