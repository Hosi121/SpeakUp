package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PROGRESS struct {
	ent.Schema
}

func (PROGRESS) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("login_days").
			Default(0),
		field.Int("consecutive_participants").
			Default(0),
		field.Int("consecutive_login_days").
			Default(0),
	}
}

func (PROGRESS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("recorded", USERS.Type).
			Ref("records").
			Unique(),
	}
}
