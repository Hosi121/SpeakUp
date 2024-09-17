package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// USERS holds the schema definition for the USERS entity.
type CHATS struct {
	ent.Schema
}

// Fields of the USERS.
func (CHATS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("friend_id"),
		field.String("message").
			MaxLen(255),
		field.Bool("is_recieved").
			Default(false),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the USERS.
func (CHATS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("had", FRIENDS.Type).
			Ref("has").
			Unique(),
	}
}
