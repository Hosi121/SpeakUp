package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ACHIEVEMENTS struct {
	ent.Schema
}

func (ACHIEVEMENTS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("trophy_id"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (ACHIEVEMENTS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("granted", USERS.Type).
			Ref("acquires").
			Unique(),
		edge.To("refers", TROPHIES.Type).
			Unique(),
	}
}
