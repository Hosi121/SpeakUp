package schema

import "entgo.io/ent"

// USERS holds the schema definition for the USERS entity.
type USERS struct {
	ent.Schema
}

// Fields of the USERS.
func (USERS) Fields() []ent.Field {
	return nil
}

// Edges of the USERS.
func (USERS) Edges() []ent.Edge {
	return nil
}
