package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FRIENDS holds the schema definition for the FRIENDS entity.
type FRIENDS struct {
	ent.Schema
}

// Fields of the FRIENDS.
func (FRIENDS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("friend_id").
			Unique(),
		field.Int("user_id"),
		field.Int("target_user_id"),
		field.Enum("status").
			Values("PENDING", "FRIEND", "BLOCKED"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the FRIENDS.
func (FRIENDS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("connects", USERS.Type).
			Ref("connects"),
	}
}
