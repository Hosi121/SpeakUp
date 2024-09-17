package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// MATCHINGS holds the schema definition for the MATCHINGS entity.
type SESSIONS struct {
	ent.Schema
}

// Fields of the MATCHINGS.
func (SESSIONS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("matched_user_id"),
		field.Int("record_id"),
		field.Time("matched_at").
			Default(time.Now()),
		field.Enum("status").
			Values("MATCHED", "PROCCESSING", "FINISHED"),
	}
}

// Edges of the MATCHINGS.
func (SESSIONS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("had", EVENT_RECORDS.Type).
			Ref("has").
			Unique(),
		edge.To("makes", CALLS.Type).
			Unique(),
	}
}
