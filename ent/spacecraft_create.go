// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"spacecraft/ent/spacecraft"
	"spacecraft/ent/spacecraftarmament"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpacecraftCreate is the builder for creating a Spacecraft entity.
type SpacecraftCreate struct {
	config
	mutation *SpacecraftMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (sc *SpacecraftCreate) SetName(s string) *SpacecraftCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetClass sets the "class" field.
func (sc *SpacecraftCreate) SetClass(s string) *SpacecraftCreate {
	sc.mutation.SetClass(s)
	return sc
}

// SetCrew sets the "crew" field.
func (sc *SpacecraftCreate) SetCrew(i int) *SpacecraftCreate {
	sc.mutation.SetCrew(i)
	return sc
}

// SetImage sets the "image" field.
func (sc *SpacecraftCreate) SetImage(s string) *SpacecraftCreate {
	sc.mutation.SetImage(s)
	return sc
}

// SetValue sets the "value" field.
func (sc *SpacecraftCreate) SetValue(f float64) *SpacecraftCreate {
	sc.mutation.SetValue(f)
	return sc
}

// SetStatus sets the "status" field.
func (sc *SpacecraftCreate) SetStatus(s string) *SpacecraftCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SpacecraftCreate) SetCreatedAt(t time.Time) *SpacecraftCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SpacecraftCreate) SetNillableCreatedAt(t *time.Time) *SpacecraftCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SpacecraftCreate) SetUpdatedAt(t time.Time) *SpacecraftCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SpacecraftCreate) SetNillableUpdatedAt(t *time.Time) *SpacecraftCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SpacecraftCreate) SetDeletedAt(t time.Time) *SpacecraftCreate {
	sc.mutation.SetDeletedAt(t)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SpacecraftCreate) SetNillableDeletedAt(t *time.Time) *SpacecraftCreate {
	if t != nil {
		sc.SetDeletedAt(*t)
	}
	return sc
}

// AddArmamentIDs adds the "armaments" edge to the SpacecraftArmament entity by IDs.
func (sc *SpacecraftCreate) AddArmamentIDs(ids ...int) *SpacecraftCreate {
	sc.mutation.AddArmamentIDs(ids...)
	return sc
}

// AddArmaments adds the "armaments" edges to the SpacecraftArmament entity.
func (sc *SpacecraftCreate) AddArmaments(s ...*SpacecraftArmament) *SpacecraftCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddArmamentIDs(ids...)
}

// Mutation returns the SpacecraftMutation object of the builder.
func (sc *SpacecraftCreate) Mutation() *SpacecraftMutation {
	return sc.mutation
}

// Save creates the Spacecraft in the database.
func (sc *SpacecraftCreate) Save(ctx context.Context) (*Spacecraft, error) {
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SpacecraftCreate) SaveX(ctx context.Context) *Spacecraft {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SpacecraftCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SpacecraftCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SpacecraftCreate) check() error {
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Spacecraft.name"`)}
	}
	if _, ok := sc.mutation.Class(); !ok {
		return &ValidationError{Name: "class", err: errors.New(`ent: missing required field "Spacecraft.class"`)}
	}
	if _, ok := sc.mutation.Crew(); !ok {
		return &ValidationError{Name: "crew", err: errors.New(`ent: missing required field "Spacecraft.crew"`)}
	}
	if _, ok := sc.mutation.Image(); !ok {
		return &ValidationError{Name: "image", err: errors.New(`ent: missing required field "Spacecraft.image"`)}
	}
	if _, ok := sc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Spacecraft.value"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Spacecraft.status"`)}
	}
	return nil
}

func (sc *SpacecraftCreate) sqlSave(ctx context.Context) (*Spacecraft, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SpacecraftCreate) createSpec() (*Spacecraft, *sqlgraph.CreateSpec) {
	var (
		_node = &Spacecraft{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(spacecraft.Table, sqlgraph.NewFieldSpec(spacecraft.FieldID, field.TypeInt))
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(spacecraft.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Class(); ok {
		_spec.SetField(spacecraft.FieldClass, field.TypeString, value)
		_node.Class = value
	}
	if value, ok := sc.mutation.Crew(); ok {
		_spec.SetField(spacecraft.FieldCrew, field.TypeInt, value)
		_node.Crew = value
	}
	if value, ok := sc.mutation.Image(); ok {
		_spec.SetField(spacecraft.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(spacecraft.FieldValue, field.TypeFloat64, value)
		_node.Value = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(spacecraft.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(spacecraft.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(spacecraft.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.SetField(spacecraft.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if nodes := sc.mutation.ArmamentsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Spacecraft.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SpacecraftUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (sc *SpacecraftCreate) OnConflict(opts ...sql.ConflictOption) *SpacecraftUpsertOne {
	sc.conflict = opts
	return &SpacecraftUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Spacecraft.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SpacecraftCreate) OnConflictColumns(columns ...string) *SpacecraftUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SpacecraftUpsertOne{
		create: sc,
	}
}

type (
	// SpacecraftUpsertOne is the builder for "upsert"-ing
	//  one Spacecraft node.
	SpacecraftUpsertOne struct {
		create *SpacecraftCreate
	}

	// SpacecraftUpsert is the "OnConflict" setter.
	SpacecraftUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *SpacecraftUpsert) SetName(v string) *SpacecraftUpsert {
	u.Set(spacecraft.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateName() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldName)
	return u
}

// SetClass sets the "class" field.
func (u *SpacecraftUpsert) SetClass(v string) *SpacecraftUpsert {
	u.Set(spacecraft.FieldClass, v)
	return u
}

// UpdateClass sets the "class" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateClass() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldClass)
	return u
}

// SetCrew sets the "crew" field.
func (u *SpacecraftUpsert) SetCrew(v int) *SpacecraftUpsert {
	u.Set(spacecraft.FieldCrew, v)
	return u
}

// UpdateCrew sets the "crew" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateCrew() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldCrew)
	return u
}

// AddCrew adds v to the "crew" field.
func (u *SpacecraftUpsert) AddCrew(v int) *SpacecraftUpsert {
	u.Add(spacecraft.FieldCrew, v)
	return u
}

// SetImage sets the "image" field.
func (u *SpacecraftUpsert) SetImage(v string) *SpacecraftUpsert {
	u.Set(spacecraft.FieldImage, v)
	return u
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateImage() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldImage)
	return u
}

// SetValue sets the "value" field.
func (u *SpacecraftUpsert) SetValue(v float64) *SpacecraftUpsert {
	u.Set(spacecraft.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateValue() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldValue)
	return u
}

// AddValue adds v to the "value" field.
func (u *SpacecraftUpsert) AddValue(v float64) *SpacecraftUpsert {
	u.Add(spacecraft.FieldValue, v)
	return u
}

// SetStatus sets the "status" field.
func (u *SpacecraftUpsert) SetStatus(v string) *SpacecraftUpsert {
	u.Set(spacecraft.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateStatus() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldStatus)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SpacecraftUpsert) SetCreatedAt(v time.Time) *SpacecraftUpsert {
	u.Set(spacecraft.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateCreatedAt() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldCreatedAt)
	return u
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *SpacecraftUpsert) ClearCreatedAt() *SpacecraftUpsert {
	u.SetNull(spacecraft.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SpacecraftUpsert) SetUpdatedAt(v time.Time) *SpacecraftUpsert {
	u.Set(spacecraft.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateUpdatedAt() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldUpdatedAt)
	return u
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SpacecraftUpsert) ClearUpdatedAt() *SpacecraftUpsert {
	u.SetNull(spacecraft.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SpacecraftUpsert) SetDeletedAt(v time.Time) *SpacecraftUpsert {
	u.Set(spacecraft.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SpacecraftUpsert) UpdateDeletedAt() *SpacecraftUpsert {
	u.SetExcluded(spacecraft.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SpacecraftUpsert) ClearDeletedAt() *SpacecraftUpsert {
	u.SetNull(spacecraft.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Spacecraft.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SpacecraftUpsertOne) UpdateNewValues() *SpacecraftUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Spacecraft.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SpacecraftUpsertOne) Ignore() *SpacecraftUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SpacecraftUpsertOne) DoNothing() *SpacecraftUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SpacecraftCreate.OnConflict
// documentation for more info.
func (u *SpacecraftUpsertOne) Update(set func(*SpacecraftUpsert)) *SpacecraftUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SpacecraftUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *SpacecraftUpsertOne) SetName(v string) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateName() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateName()
	})
}

// SetClass sets the "class" field.
func (u *SpacecraftUpsertOne) SetClass(v string) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetClass(v)
	})
}

// UpdateClass sets the "class" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateClass() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateClass()
	})
}

// SetCrew sets the "crew" field.
func (u *SpacecraftUpsertOne) SetCrew(v int) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetCrew(v)
	})
}

// AddCrew adds v to the "crew" field.
func (u *SpacecraftUpsertOne) AddCrew(v int) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.AddCrew(v)
	})
}

// UpdateCrew sets the "crew" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateCrew() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateCrew()
	})
}

// SetImage sets the "image" field.
func (u *SpacecraftUpsertOne) SetImage(v string) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateImage() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateImage()
	})
}

// SetValue sets the "value" field.
func (u *SpacecraftUpsertOne) SetValue(v float64) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *SpacecraftUpsertOne) AddValue(v float64) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateValue() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateValue()
	})
}

// SetStatus sets the "status" field.
func (u *SpacecraftUpsertOne) SetStatus(v string) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateStatus() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateStatus()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SpacecraftUpsertOne) SetCreatedAt(v time.Time) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateCreatedAt() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *SpacecraftUpsertOne) ClearCreatedAt() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SpacecraftUpsertOne) SetUpdatedAt(v time.Time) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateUpdatedAt() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SpacecraftUpsertOne) ClearUpdatedAt() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SpacecraftUpsertOne) SetDeletedAt(v time.Time) *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SpacecraftUpsertOne) UpdateDeletedAt() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SpacecraftUpsertOne) ClearDeletedAt() *SpacecraftUpsertOne {
	return u.Update(func(s *SpacecraftUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SpacecraftUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SpacecraftCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SpacecraftUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SpacecraftUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SpacecraftUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SpacecraftCreateBulk is the builder for creating many Spacecraft entities in bulk.
type SpacecraftCreateBulk struct {
	config
	err      error
	builders []*SpacecraftCreate
	conflict []sql.ConflictOption
}

// Save creates the Spacecraft entities in the database.
func (scb *SpacecraftCreateBulk) Save(ctx context.Context) ([]*Spacecraft, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Spacecraft, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SpacecraftMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SpacecraftCreateBulk) SaveX(ctx context.Context) []*Spacecraft {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SpacecraftCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SpacecraftCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Spacecraft.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SpacecraftUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (scb *SpacecraftCreateBulk) OnConflict(opts ...sql.ConflictOption) *SpacecraftUpsertBulk {
	scb.conflict = opts
	return &SpacecraftUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Spacecraft.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SpacecraftCreateBulk) OnConflictColumns(columns ...string) *SpacecraftUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SpacecraftUpsertBulk{
		create: scb,
	}
}

// SpacecraftUpsertBulk is the builder for "upsert"-ing
// a bulk of Spacecraft nodes.
type SpacecraftUpsertBulk struct {
	create *SpacecraftCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Spacecraft.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *SpacecraftUpsertBulk) UpdateNewValues() *SpacecraftUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Spacecraft.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SpacecraftUpsertBulk) Ignore() *SpacecraftUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SpacecraftUpsertBulk) DoNothing() *SpacecraftUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SpacecraftCreateBulk.OnConflict
// documentation for more info.
func (u *SpacecraftUpsertBulk) Update(set func(*SpacecraftUpsert)) *SpacecraftUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SpacecraftUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *SpacecraftUpsertBulk) SetName(v string) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateName() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateName()
	})
}

// SetClass sets the "class" field.
func (u *SpacecraftUpsertBulk) SetClass(v string) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetClass(v)
	})
}

// UpdateClass sets the "class" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateClass() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateClass()
	})
}

// SetCrew sets the "crew" field.
func (u *SpacecraftUpsertBulk) SetCrew(v int) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetCrew(v)
	})
}

// AddCrew adds v to the "crew" field.
func (u *SpacecraftUpsertBulk) AddCrew(v int) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.AddCrew(v)
	})
}

// UpdateCrew sets the "crew" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateCrew() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateCrew()
	})
}

// SetImage sets the "image" field.
func (u *SpacecraftUpsertBulk) SetImage(v string) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetImage(v)
	})
}

// UpdateImage sets the "image" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateImage() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateImage()
	})
}

// SetValue sets the "value" field.
func (u *SpacecraftUpsertBulk) SetValue(v float64) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetValue(v)
	})
}

// AddValue adds v to the "value" field.
func (u *SpacecraftUpsertBulk) AddValue(v float64) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.AddValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateValue() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateValue()
	})
}

// SetStatus sets the "status" field.
func (u *SpacecraftUpsertBulk) SetStatus(v string) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateStatus() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateStatus()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *SpacecraftUpsertBulk) SetCreatedAt(v time.Time) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateCreatedAt() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateCreatedAt()
	})
}

// ClearCreatedAt clears the value of the "created_at" field.
func (u *SpacecraftUpsertBulk) ClearCreatedAt() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.ClearCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SpacecraftUpsertBulk) SetUpdatedAt(v time.Time) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateUpdatedAt() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateUpdatedAt()
	})
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (u *SpacecraftUpsertBulk) ClearUpdatedAt() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.ClearUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SpacecraftUpsertBulk) SetDeletedAt(v time.Time) *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SpacecraftUpsertBulk) UpdateDeletedAt() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *SpacecraftUpsertBulk) ClearDeletedAt() *SpacecraftUpsertBulk {
	return u.Update(func(s *SpacecraftUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *SpacecraftUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SpacecraftCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SpacecraftCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SpacecraftUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
