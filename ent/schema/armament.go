package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Armament holds the schema definition for the Armament entity.
type Armament struct {
	ent.Schema
}

// Fields of the Armament.
func (Armament) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Int("qty"),
	}
}

// Edges of the Armament.
func (Armament) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Spacecraft", Spacecraft.Type).
			Ref("armaments"),
	}
}
