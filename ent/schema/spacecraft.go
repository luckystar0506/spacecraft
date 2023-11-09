package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Time("created_at").Optional().Annotations(entsql.Default("CURRENT_TIMESTAMP")),
		field.Time("updated_at").Optional().Annotations(entsql.Default("CURRENT_TIMESTAMP")).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Spacecraft.
func (Spacecraft) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("armaments", SpacecraftArmament.Type),
	}
}
