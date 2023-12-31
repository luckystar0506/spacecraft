// Code generated by ent, DO NOT EDIT.

package spacecraftarmament

import (
	"spacecraft/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldLTE(FieldID, id))
}

// Qty applies equality check predicate on the "qty" field. It's identical to QtyEQ.
func Qty(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldQty, v))
}

// SpacecraftID applies equality check predicate on the "spacecraft_id" field. It's identical to SpacecraftIDEQ.
func SpacecraftID(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldSpacecraftID, v))
}

// ArmamentID applies equality check predicate on the "armament_id" field. It's identical to ArmamentIDEQ.
func ArmamentID(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldArmamentID, v))
}

// QtyEQ applies the EQ predicate on the "qty" field.
func QtyEQ(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldQty, v))
}

// QtyNEQ applies the NEQ predicate on the "qty" field.
func QtyNEQ(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNEQ(FieldQty, v))
}

// QtyIn applies the In predicate on the "qty" field.
func QtyIn(vs ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldIn(FieldQty, vs...))
}

// QtyNotIn applies the NotIn predicate on the "qty" field.
func QtyNotIn(vs ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNotIn(FieldQty, vs...))
}

// QtyGT applies the GT predicate on the "qty" field.
func QtyGT(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldGT(FieldQty, v))
}

// QtyGTE applies the GTE predicate on the "qty" field.
func QtyGTE(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldGTE(FieldQty, v))
}

// QtyLT applies the LT predicate on the "qty" field.
func QtyLT(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldLT(FieldQty, v))
}

// QtyLTE applies the LTE predicate on the "qty" field.
func QtyLTE(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldLTE(FieldQty, v))
}

// SpacecraftIDEQ applies the EQ predicate on the "spacecraft_id" field.
func SpacecraftIDEQ(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldSpacecraftID, v))
}

// SpacecraftIDNEQ applies the NEQ predicate on the "spacecraft_id" field.
func SpacecraftIDNEQ(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNEQ(FieldSpacecraftID, v))
}

// SpacecraftIDIn applies the In predicate on the "spacecraft_id" field.
func SpacecraftIDIn(vs ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldIn(FieldSpacecraftID, vs...))
}

// SpacecraftIDNotIn applies the NotIn predicate on the "spacecraft_id" field.
func SpacecraftIDNotIn(vs ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNotIn(FieldSpacecraftID, vs...))
}

// ArmamentIDEQ applies the EQ predicate on the "armament_id" field.
func ArmamentIDEQ(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldEQ(FieldArmamentID, v))
}

// ArmamentIDNEQ applies the NEQ predicate on the "armament_id" field.
func ArmamentIDNEQ(v int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNEQ(FieldArmamentID, v))
}

// ArmamentIDIn applies the In predicate on the "armament_id" field.
func ArmamentIDIn(vs ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldIn(FieldArmamentID, vs...))
}

// ArmamentIDNotIn applies the NotIn predicate on the "armament_id" field.
func ArmamentIDNotIn(vs ...int) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.FieldNotIn(FieldArmamentID, vs...))
}

// HasSpacecraft applies the HasEdge predicate on the "spacecraft" edge.
func HasSpacecraft() predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SpacecraftTable, SpacecraftColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpacecraftWith applies the HasEdge predicate on the "spacecraft" edge with a given conditions (other predicates).
func HasSpacecraftWith(preds ...predicate.Spacecraft) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(func(s *sql.Selector) {
		step := newSpacecraftStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasArmament applies the HasEdge predicate on the "armament" edge.
func HasArmament() predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ArmamentTable, ArmamentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArmamentWith applies the HasEdge predicate on the "armament" edge with a given conditions (other predicates).
func HasArmamentWith(preds ...predicate.Armament) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(func(s *sql.Selector) {
		step := newArmamentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SpacecraftArmament) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SpacecraftArmament) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.SpacecraftArmament) predicate.SpacecraftArmament {
	return predicate.SpacecraftArmament(sql.NotPredicates(p))
}
