package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Spacecraft holds the schema definition for the Spacecraft entity.
type Spacecraft struct {
	ent.Schema
}

// Fields of the Spacecraft.
func (Spacecraft) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("class"),
		field.Int("crew"),
		field.String("image"),
		field.Float("value"),
		field.String("status"),
	}
}

// Edges of the Spacecraft.
func (Spacecraft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("armaments", Armament.Type),
	}
}
