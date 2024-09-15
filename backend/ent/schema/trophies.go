package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TROPHIES struct {
	ent.Schema
}

func (TROPHIES) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("contents"),
		field.String("requirement"),
	}
}

func (TROPHIES) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("refered", ACHIEVEMENTS.Type).
			Ref("refers"),
	}
}
