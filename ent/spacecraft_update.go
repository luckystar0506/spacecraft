// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"spacecraft/ent/predicate"
	"spacecraft/ent/spacecraft"
	"spacecraft/ent/spacecraftarmament"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpacecraftUpdate is the builder for updating Spacecraft entities.
type SpacecraftUpdate struct {
	config
	hooks    []Hook
	mutation *SpacecraftMutation
}

// Where appends a list predicates to the SpacecraftUpdate builder.
func (su *SpacecraftUpdate) Where(ps ...predicate.Spacecraft) *SpacecraftUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetName sets the "name" field.
func (su *SpacecraftUpdate) SetName(s string) *SpacecraftUpdate {
	su.mutation.SetName(s)
	return su
}

// SetClass sets the "class" field.
func (su *SpacecraftUpdate) SetClass(s string) *SpacecraftUpdate {
	su.mutation.SetClass(s)
	return su
}

// SetCrew sets the "crew" field.
func (su *SpacecraftUpdate) SetCrew(i int) *SpacecraftUpdate {
	su.mutation.ResetCrew()
	su.mutation.SetCrew(i)
	return su
}

// AddCrew adds i to the "crew" field.
func (su *SpacecraftUpdate) AddCrew(i int) *SpacecraftUpdate {
	su.mutation.AddCrew(i)
	return su
}

// SetImage sets the "image" field.
func (su *SpacecraftUpdate) SetImage(s string) *SpacecraftUpdate {
	su.mutation.SetImage(s)
	return su
}

// SetValue sets the "value" field.
func (su *SpacecraftUpdate) SetValue(f float64) *SpacecraftUpdate {
	su.mutation.ResetValue()
	su.mutation.SetValue(f)
	return su
}

// AddValue adds f to the "value" field.
func (su *SpacecraftUpdate) AddValue(f float64) *SpacecraftUpdate {
	su.mutation.AddValue(f)
	return su
}

// SetStatus sets the "status" field.
func (su *SpacecraftUpdate) SetStatus(s string) *SpacecraftUpdate {
	su.mutation.SetStatus(s)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SpacecraftUpdate) SetCreatedAt(t time.Time) *SpacecraftUpdate {
	su.mutation.SetCreatedAt(t)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SpacecraftUpdate) SetNillableCreatedAt(t *time.Time) *SpacecraftUpdate {
	if t != nil {
		su.SetCreatedAt(*t)
	}
	return su
}

// ClearCreatedAt clears the value of the "created_at" field.
func (su *SpacecraftUpdate) ClearCreatedAt() *SpacecraftUpdate {
	su.mutation.ClearCreatedAt()
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SpacecraftUpdate) SetUpdatedAt(t time.Time) *SpacecraftUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (su *SpacecraftUpdate) ClearUpdatedAt() *SpacecraftUpdate {
	su.mutation.ClearUpdatedAt()
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *SpacecraftUpdate) SetDeletedAt(t time.Time) *SpacecraftUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *SpacecraftUpdate) SetNillableDeletedAt(t *time.Time) *SpacecraftUpdate {
	if t != nil {
		su.SetDeletedAt(*t)
	}
	return su
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (su *SpacecraftUpdate) ClearDeletedAt() *SpacecraftUpdate {
	su.mutation.ClearDeletedAt()
	return su
}

// AddArmamentIDs adds the "armaments" edge to the SpacecraftArmament entity by IDs.
func (su *SpacecraftUpdate) AddArmamentIDs(ids ...int) *SpacecraftUpdate {
	su.mutation.AddArmamentIDs(ids...)
	return su
}

// AddArmaments adds the "armaments" edges to the SpacecraftArmament entity.
func (su *SpacecraftUpdate) AddArmaments(s ...*SpacecraftArmament) *SpacecraftUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddArmamentIDs(ids...)
}

// Mutation returns the SpacecraftMutation object of the builder.
func (su *SpacecraftUpdate) Mutation() *SpacecraftMutation {
	return su.mutation
}

// ClearArmaments clears all "armaments" edges to the SpacecraftArmament entity.
func (su *SpacecraftUpdate) ClearArmaments() *SpacecraftUpdate {
	su.mutation.ClearArmaments()
	return su
}

// RemoveArmamentIDs removes the "armaments" edge to SpacecraftArmament entities by IDs.
func (su *SpacecraftUpdate) RemoveArmamentIDs(ids ...int) *SpacecraftUpdate {
	su.mutation.RemoveArmamentIDs(ids...)
	return su
}

// RemoveArmaments removes "armaments" edges to SpacecraftArmament entities.
func (su *SpacecraftUpdate) RemoveArmaments(s ...*SpacecraftArmament) *SpacecraftUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveArmamentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SpacecraftUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SpacecraftUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SpacecraftUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SpacecraftUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SpacecraftUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok && !su.mutation.UpdatedAtCleared() {
		v := spacecraft.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *SpacecraftUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(spacecraft.Table, spacecraft.Columns, sqlgraph.NewFieldSpec(spacecraft.FieldID, field.TypeInt))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(spacecraft.FieldName, field.TypeString, value)
	}
	if value, ok := su.mutation.Class(); ok {
		_spec.SetField(spacecraft.FieldClass, field.TypeString, value)
	}
	if value, ok := su.mutation.Crew(); ok {
		_spec.SetField(spacecraft.FieldCrew, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedCrew(); ok {
		_spec.AddField(spacecraft.FieldCrew, field.TypeInt, value)
	}
	if value, ok := su.mutation.Image(); ok {
		_spec.SetField(spacecraft.FieldImage, field.TypeString, value)
	}
	if value, ok := su.mutation.Value(); ok {
		_spec.SetField(spacecraft.FieldValue, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.AddedValue(); ok {
		_spec.AddField(spacecraft.FieldValue, field.TypeFloat64, value)
	}
	if value, ok := su.mutation.Status(); ok {
		_spec.SetField(spacecraft.FieldStatus, field.TypeString, value)
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.SetField(spacecraft.FieldCreatedAt, field.TypeTime, value)
	}
	if su.mutation.CreatedAtCleared() {
		_spec.ClearField(spacecraft.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(spacecraft.FieldUpdatedAt, field.TypeTime, value)
	}
	if su.mutation.UpdatedAtCleared() {
		_spec.ClearField(spacecraft.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.SetField(spacecraft.FieldDeletedAt, field.TypeTime, value)
	}
	if su.mutation.DeletedAtCleared() {
		_spec.ClearField(spacecraft.FieldDeletedAt, field.TypeTime)
	}
	if su.mutation.ArmamentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spacecraft.ArmamentsTable,
			Columns: []string{spacecraft.ArmamentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedArmamentsIDs(); len(nodes) > 0 && !su.mutation.ArmamentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spacecraft.ArmamentsTable,
			Columns: []string{spacecraft.ArmamentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ArmamentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spacecraft.ArmamentsTable,
			Columns: []string{spacecraft.ArmamentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spacecraft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SpacecraftUpdateOne is the builder for updating a single Spacecraft entity.
type SpacecraftUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SpacecraftMutation
}

// SetName sets the "name" field.
func (suo *SpacecraftUpdateOne) SetName(s string) *SpacecraftUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetClass sets the "class" field.
func (suo *SpacecraftUpdateOne) SetClass(s string) *SpacecraftUpdateOne {
	suo.mutation.SetClass(s)
	return suo
}

// SetCrew sets the "crew" field.
func (suo *SpacecraftUpdateOne) SetCrew(i int) *SpacecraftUpdateOne {
	suo.mutation.ResetCrew()
	suo.mutation.SetCrew(i)
	return suo
}

// AddCrew adds i to the "crew" field.
func (suo *SpacecraftUpdateOne) AddCrew(i int) *SpacecraftUpdateOne {
	suo.mutation.AddCrew(i)
	return suo
}

// SetImage sets the "image" field.
func (suo *SpacecraftUpdateOne) SetImage(s string) *SpacecraftUpdateOne {
	suo.mutation.SetImage(s)
	return suo
}

// SetValue sets the "value" field.
func (suo *SpacecraftUpdateOne) SetValue(f float64) *SpacecraftUpdateOne {
	suo.mutation.ResetValue()
	suo.mutation.SetValue(f)
	return suo
}

// AddValue adds f to the "value" field.
func (suo *SpacecraftUpdateOne) AddValue(f float64) *SpacecraftUpdateOne {
	suo.mutation.AddValue(f)
	return suo
}

// SetStatus sets the "status" field.
func (suo *SpacecraftUpdateOne) SetStatus(s string) *SpacecraftUpdateOne {
	suo.mutation.SetStatus(s)
	return suo
}

// SetCreatedAt sets the "created_at" field.
func (suo *SpacecraftUpdateOne) SetCreatedAt(t time.Time) *SpacecraftUpdateOne {
	suo.mutation.SetCreatedAt(t)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SpacecraftUpdateOne) SetNillableCreatedAt(t *time.Time) *SpacecraftUpdateOne {
	if t != nil {
		suo.SetCreatedAt(*t)
	}
	return suo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (suo *SpacecraftUpdateOne) ClearCreatedAt() *SpacecraftUpdateOne {
	suo.mutation.ClearCreatedAt()
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SpacecraftUpdateOne) SetUpdatedAt(t time.Time) *SpacecraftUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (suo *SpacecraftUpdateOne) ClearUpdatedAt() *SpacecraftUpdateOne {
	suo.mutation.ClearUpdatedAt()
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *SpacecraftUpdateOne) SetDeletedAt(t time.Time) *SpacecraftUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *SpacecraftUpdateOne) SetNillableDeletedAt(t *time.Time) *SpacecraftUpdateOne {
	if t != nil {
		suo.SetDeletedAt(*t)
	}
	return suo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (suo *SpacecraftUpdateOne) ClearDeletedAt() *SpacecraftUpdateOne {
	suo.mutation.ClearDeletedAt()
	return suo
}

// AddArmamentIDs adds the "armaments" edge to the SpacecraftArmament entity by IDs.
func (suo *SpacecraftUpdateOne) AddArmamentIDs(ids ...int) *SpacecraftUpdateOne {
	suo.mutation.AddArmamentIDs(ids...)
	return suo
}

// AddArmaments adds the "armaments" edges to the SpacecraftArmament entity.
func (suo *SpacecraftUpdateOne) AddArmaments(s ...*SpacecraftArmament) *SpacecraftUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddArmamentIDs(ids...)
}

// Mutation returns the SpacecraftMutation object of the builder.
func (suo *SpacecraftUpdateOne) Mutation() *SpacecraftMutation {
	return suo.mutation
}

// ClearArmaments clears all "armaments" edges to the SpacecraftArmament entity.
func (suo *SpacecraftUpdateOne) ClearArmaments() *SpacecraftUpdateOne {
	suo.mutation.ClearArmaments()
	return suo
}

// RemoveArmamentIDs removes the "armaments" edge to SpacecraftArmament entities by IDs.
func (suo *SpacecraftUpdateOne) RemoveArmamentIDs(ids ...int) *SpacecraftUpdateOne {
	suo.mutation.RemoveArmamentIDs(ids...)
	return suo
}

// RemoveArmaments removes "armaments" edges to SpacecraftArmament entities.
func (suo *SpacecraftUpdateOne) RemoveArmaments(s ...*SpacecraftArmament) *SpacecraftUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveArmamentIDs(ids...)
}

// Where appends a list predicates to the SpacecraftUpdate builder.
func (suo *SpacecraftUpdateOne) Where(ps ...predicate.Spacecraft) *SpacecraftUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SpacecraftUpdateOne) Select(field string, fields ...string) *SpacecraftUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Spacecraft entity.
func (suo *SpacecraftUpdateOne) Save(ctx context.Context) (*Spacecraft, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SpacecraftUpdateOne) SaveX(ctx context.Context) *Spacecraft {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SpacecraftUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SpacecraftUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SpacecraftUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok && !suo.mutation.UpdatedAtCleared() {
		v := spacecraft.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *SpacecraftUpdateOne) sqlSave(ctx context.Context) (_node *Spacecraft, err error) {
	_spec := sqlgraph.NewUpdateSpec(spacecraft.Table, spacecraft.Columns, sqlgraph.NewFieldSpec(spacecraft.FieldID, field.TypeInt))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Spacecraft.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, spacecraft.FieldID)
		for _, f := range fields {
			if !spacecraft.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != spacecraft.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(spacecraft.FieldName, field.TypeString, value)
	}
	if value, ok := suo.mutation.Class(); ok {
		_spec.SetField(spacecraft.FieldClass, field.TypeString, value)
	}
	if value, ok := suo.mutation.Crew(); ok {
		_spec.SetField(spacecraft.FieldCrew, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedCrew(); ok {
		_spec.AddField(spacecraft.FieldCrew, field.TypeInt, value)
	}
	if value, ok := suo.mutation.Image(); ok {
		_spec.SetField(spacecraft.FieldImage, field.TypeString, value)
	}
	if value, ok := suo.mutation.Value(); ok {
		_spec.SetField(spacecraft.FieldValue, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.AddedValue(); ok {
		_spec.AddField(spacecraft.FieldValue, field.TypeFloat64, value)
	}
	if value, ok := suo.mutation.Status(); ok {
		_spec.SetField(spacecraft.FieldStatus, field.TypeString, value)
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.SetField(spacecraft.FieldCreatedAt, field.TypeTime, value)
	}
	if suo.mutation.CreatedAtCleared() {
		_spec.ClearField(spacecraft.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(spacecraft.FieldUpdatedAt, field.TypeTime, value)
	}
	if suo.mutation.UpdatedAtCleared() {
		_spec.ClearField(spacecraft.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.SetField(spacecraft.FieldDeletedAt, field.TypeTime, value)
	}
	if suo.mutation.DeletedAtCleared() {
		_spec.ClearField(spacecraft.FieldDeletedAt, field.TypeTime)
	}
	if suo.mutation.ArmamentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spacecraft.ArmamentsTable,
			Columns: []string{spacecraft.ArmamentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedArmamentsIDs(); len(nodes) > 0 && !suo.mutation.ArmamentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spacecraft.ArmamentsTable,
			Columns: []string{spacecraft.ArmamentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ArmamentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   spacecraft.ArmamentsTable,
			Columns: []string{spacecraft.ArmamentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Spacecraft{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{spacecraft.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
