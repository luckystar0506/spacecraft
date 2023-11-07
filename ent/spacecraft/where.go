// Code generated by ent, DO NOT EDIT.

package spacecraft

import (
	"spacecraft/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldName, v))
}

// Class applies equality check predicate on the "class" field. It's identical to ClassEQ.
func Class(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldClass, v))
}

// Crew applies equality check predicate on the "crew" field. It's identical to CrewEQ.
func Crew(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldCrew, v))
}

// Image applies equality check predicate on the "image" field. It's identical to ImageEQ.
func Image(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldImage, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldValue, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldStatus, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContainsFold(FieldName, v))
}

// ClassEQ applies the EQ predicate on the "class" field.
func ClassEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldClass, v))
}

// ClassNEQ applies the NEQ predicate on the "class" field.
func ClassNEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldClass, v))
}

// ClassIn applies the In predicate on the "class" field.
func ClassIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldClass, vs...))
}

// ClassNotIn applies the NotIn predicate on the "class" field.
func ClassNotIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldClass, vs...))
}

// ClassGT applies the GT predicate on the "class" field.
func ClassGT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldClass, v))
}

// ClassGTE applies the GTE predicate on the "class" field.
func ClassGTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldClass, v))
}

// ClassLT applies the LT predicate on the "class" field.
func ClassLT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldClass, v))
}

// ClassLTE applies the LTE predicate on the "class" field.
func ClassLTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldClass, v))
}

// ClassContains applies the Contains predicate on the "class" field.
func ClassContains(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContains(FieldClass, v))
}

// ClassHasPrefix applies the HasPrefix predicate on the "class" field.
func ClassHasPrefix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasPrefix(FieldClass, v))
}

// ClassHasSuffix applies the HasSuffix predicate on the "class" field.
func ClassHasSuffix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasSuffix(FieldClass, v))
}

// ClassEqualFold applies the EqualFold predicate on the "class" field.
func ClassEqualFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEqualFold(FieldClass, v))
}

// ClassContainsFold applies the ContainsFold predicate on the "class" field.
func ClassContainsFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContainsFold(FieldClass, v))
}

// CrewEQ applies the EQ predicate on the "crew" field.
func CrewEQ(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldCrew, v))
}

// CrewNEQ applies the NEQ predicate on the "crew" field.
func CrewNEQ(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldCrew, v))
}

// CrewIn applies the In predicate on the "crew" field.
func CrewIn(vs ...int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldCrew, vs...))
}

// CrewNotIn applies the NotIn predicate on the "crew" field.
func CrewNotIn(vs ...int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldCrew, vs...))
}

// CrewGT applies the GT predicate on the "crew" field.
func CrewGT(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldCrew, v))
}

// CrewGTE applies the GTE predicate on the "crew" field.
func CrewGTE(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldCrew, v))
}

// CrewLT applies the LT predicate on the "crew" field.
func CrewLT(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldCrew, v))
}

// CrewLTE applies the LTE predicate on the "crew" field.
func CrewLTE(v int) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldCrew, v))
}

// ImageEQ applies the EQ predicate on the "image" field.
func ImageEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldImage, v))
}

// ImageNEQ applies the NEQ predicate on the "image" field.
func ImageNEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldImage, v))
}

// ImageIn applies the In predicate on the "image" field.
func ImageIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldImage, vs...))
}

// ImageNotIn applies the NotIn predicate on the "image" field.
func ImageNotIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldImage, vs...))
}

// ImageGT applies the GT predicate on the "image" field.
func ImageGT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldImage, v))
}

// ImageGTE applies the GTE predicate on the "image" field.
func ImageGTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldImage, v))
}

// ImageLT applies the LT predicate on the "image" field.
func ImageLT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldImage, v))
}

// ImageLTE applies the LTE predicate on the "image" field.
func ImageLTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldImage, v))
}

// ImageContains applies the Contains predicate on the "image" field.
func ImageContains(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContains(FieldImage, v))
}

// ImageHasPrefix applies the HasPrefix predicate on the "image" field.
func ImageHasPrefix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasPrefix(FieldImage, v))
}

// ImageHasSuffix applies the HasSuffix predicate on the "image" field.
func ImageHasSuffix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasSuffix(FieldImage, v))
}

// ImageEqualFold applies the EqualFold predicate on the "image" field.
func ImageEqualFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEqualFold(FieldImage, v))
}

// ImageContainsFold applies the ContainsFold predicate on the "image" field.
func ImageContainsFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContainsFold(FieldImage, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v float64) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldValue, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Spacecraft {
	return predicate.Spacecraft(sql.FieldContainsFold(FieldStatus, v))
}

// HasArmaments applies the HasEdge predicate on the "armaments" edge.
func HasArmaments() predicate.Spacecraft {
	return predicate.Spacecraft(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ArmamentsTable, ArmamentsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasArmamentsWith applies the HasEdge predicate on the "armaments" edge with a given conditions (other predicates).
func HasArmamentsWith(preds ...predicate.Armament) predicate.Spacecraft {
	return predicate.Spacecraft(func(s *sql.Selector) {
		step := newArmamentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Spacecraft) predicate.Spacecraft {
	return predicate.Spacecraft(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Spacecraft) predicate.Spacecraft {
	return predicate.Spacecraft(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Spacecraft) predicate.Spacecraft {
	return predicate.Spacecraft(sql.NotPredicates(p))
}
