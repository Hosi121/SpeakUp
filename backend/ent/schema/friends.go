package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// FRIENDS holds the schema definition for the FRIENDS entity.
type FRIENDS struct {
	ent.Schema
}

// Fields of the FRIENDS.
func (FRIENDS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Unique().
			Immutable(),
		field.Int("friend_id"),
		field.Enum("status").
			Values("PENDING", "FRIEND", "BLOCKED"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the FRIENDS.
func (FRIENDS) Edges() []ent.Edge {
	return nil
}
