package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// USERS holds the schema definition for the USERS entity.
type MEMOS struct {
	ent.Schema
}

// Fields of the USERS.
func (MEMOS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Unique(),
		field.String("memo1").
			MaxLen(255).
			Default(""),
		field.String("memo2").
			MaxLen(255).
			Default(""),
	}
}

// Edges of the USERS.
func (MEMOS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prepared", USERS.Type).
			Ref("prepares").
			Unique(),
	}
}
