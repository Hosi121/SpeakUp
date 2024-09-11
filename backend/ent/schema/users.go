package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// USERS holds the schema definition for the USERS entity.
type USERS struct {
	ent.Schema
}

// Fields of the USERS.
func (USERS) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			NotEmpty().
			MaxLen(255),
		field.String("email").
			NotEmpty().
			MaxLen(255),
		field.Text("hashed_password").
			NotEmpty(),
		field.Text("avatar_url").
			Optional(),
		field.Enum("role").
			Values("SUPERUSER", "ADMIN", "USER"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("deleted_at").
			Default(time.Now),
	}
}

// Edges of the USERS.
func (USERS) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("connects", FRIENDS.Type),
		edge.To("participates", MATCHINGS.Type),
		edge.To("makes", CALLS.Type),
		edge.To("creates", SESSIONS.Type),
	}
}
