package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// SpacecraftArmament holds the schema definition for the SpacecraftArmament entity.
type SpacecraftArmament struct {
	ent.Schema
}

// Fields of the SpacecraftArmament.
func (SpacecraftArmament) Fields() []ent.Field {
	return []ent.Field{
		field.Int("qty"),
		field.Int("spacecraft_id"),
		field.Int("armament_id"),
	}
}

// Edges of the SpacecraftArmament.
func (SpacecraftArmament) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("spacecraft", Spacecraft.Type).
			Ref("armaments").
			Field("spacecraft_id").
			Unique().
			Required(),
		edge.From("armament", Armament.Type).
			Ref("spacecrafts").
			Field("armament_id").
			Unique().
			Required(),
	}
}

func (SpacecraftArmament) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("spacecraft_id", "armament_id").Unique(),
	}
}
