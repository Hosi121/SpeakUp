package schema

import "entgo.io/ent"

// FRIENDS holds the schema definition for the FRIENDS entity.
type FRIENDS struct {
	ent.Schema
}

// Fields of the FRIENDS.
func (FRIENDS) Fields() []ent.Field {
	return nil
}

// Edges of the FRIENDS.
func (FRIENDS) Edges() []ent.Edge {
	return nil
}
