// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"spacecraft/ent/armament"
	"spacecraft/ent/predicate"
	"spacecraft/ent/spacecraft"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ArmamentQuery is the builder for querying Armament entities.
type ArmamentQuery struct {
	config
	ctx            *QueryContext
	order          []armament.OrderOption
	inters         []Interceptor
	predicates     []predicate.Armament
	withSpacecraft *SpacecraftQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArmamentQuery builder.
func (aq *ArmamentQuery) Where(ps ...predicate.Armament) *ArmamentQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ArmamentQuery) Limit(limit int) *ArmamentQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ArmamentQuery) Offset(offset int) *ArmamentQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ArmamentQuery) Unique(unique bool) *ArmamentQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ArmamentQuery) Order(o ...armament.OrderOption) *ArmamentQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QuerySpacecraft chains the current query on the "Spacecraft" edge.
func (aq *ArmamentQuery) QuerySpacecraft() *SpacecraftQuery {
	query := (&SpacecraftClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(armament.Table, armament.FieldID, selector),
			sqlgraph.To(spacecraft.Table, spacecraft.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, armament.SpacecraftTable, armament.SpacecraftPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Armament entity from the query.
// Returns a *NotFoundError when no Armament was found.
func (aq *ArmamentQuery) First(ctx context.Context) (*Armament, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{armament.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ArmamentQuery) FirstX(ctx context.Context) *Armament {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Armament ID from the query.
// Returns a *NotFoundError when no Armament ID was found.
func (aq *ArmamentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{armament.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ArmamentQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Armament entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Armament entity is found.
// Returns a *NotFoundError when no Armament entities are found.
func (aq *ArmamentQuery) Only(ctx context.Context) (*Armament, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{armament.Label}
	default:
		return nil, &NotSingularError{armament.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ArmamentQuery) OnlyX(ctx context.Context) *Armament {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Armament ID in the query.
// Returns a *NotSingularError when more than one Armament ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ArmamentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{armament.Label}
	default:
		err = &NotSingularError{armament.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ArmamentQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Armaments.
func (aq *ArmamentQuery) All(ctx context.Context) ([]*Armament, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Armament, *ArmamentQuery]()
	return withInterceptors[[]*Armament](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ArmamentQuery) AllX(ctx context.Context) []*Armament {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Armament IDs.
func (aq *ArmamentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(armament.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ArmamentQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ArmamentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ArmamentQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ArmamentQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ArmamentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ArmamentQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArmamentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ArmamentQuery) Clone() *ArmamentQuery {
	if aq == nil {
		return nil
	}
	return &ArmamentQuery{
		config:         aq.config,
		ctx:            aq.ctx.Clone(),
		order:          append([]armament.OrderOption{}, aq.order...),
		inters:         append([]Interceptor{}, aq.inters...),
		predicates:     append([]predicate.Armament{}, aq.predicates...),
		withSpacecraft: aq.withSpacecraft.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithSpacecraft tells the query-builder to eager-load the nodes that are connected to
// the "Spacecraft" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ArmamentQuery) WithSpacecraft(opts ...func(*SpacecraftQuery)) *ArmamentQuery {
	query := (&SpacecraftClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withSpacecraft = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Armament.Query().
//		GroupBy(armament.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ArmamentQuery) GroupBy(field string, fields ...string) *ArmamentGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArmamentGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = armament.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Armament.Query().
//		Select(armament.FieldTitle).
//		Scan(ctx, &v)
func (aq *ArmamentQuery) Select(fields ...string) *ArmamentSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ArmamentSelect{ArmamentQuery: aq}
	sbuild.label = armament.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArmamentSelect configured with the given aggregations.
func (aq *ArmamentQuery) Aggregate(fns ...AggregateFunc) *ArmamentSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ArmamentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !armament.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ArmamentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Armament, error) {
	var (
		nodes       = []*Armament{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withSpacecraft != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Armament).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Armament{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withSpacecraft; query != nil {
		if err := aq.loadSpacecraft(ctx, query, nodes,
			func(n *Armament) { n.Edges.Spacecraft = []*Spacecraft{} },
			func(n *Armament, e *Spacecraft) { n.Edges.Spacecraft = append(n.Edges.Spacecraft, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *ArmamentQuery) loadSpacecraft(ctx context.Context, query *SpacecraftQuery, nodes []*Armament, init func(*Armament), assign func(*Armament, *Spacecraft)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Armament)
	nids := make(map[int]map[*Armament]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(armament.SpacecraftTable)
		s.Join(joinT).On(s.C(spacecraft.FieldID), joinT.C(armament.SpacecraftPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(armament.SpacecraftPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(armament.SpacecraftPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Armament]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Spacecraft](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "Spacecraft" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (aq *ArmamentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ArmamentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(armament.Table, armament.Columns, sqlgraph.NewFieldSpec(armament.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, armament.FieldID)
		for i := range fields {
			if fields[i] != armament.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ArmamentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(armament.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = armament.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArmamentGroupBy is the group-by builder for Armament entities.
type ArmamentGroupBy struct {
	selector
	build *ArmamentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ArmamentGroupBy) Aggregate(fns ...AggregateFunc) *ArmamentGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ArmamentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArmamentQuery, *ArmamentGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ArmamentGroupBy) sqlScan(ctx context.Context, root *ArmamentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArmamentSelect is the builder for selecting fields of Armament entities.
type ArmamentSelect struct {
	*ArmamentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ArmamentSelect) Aggregate(fns ...AggregateFunc) *ArmamentSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ArmamentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArmamentQuery, *ArmamentSelect](ctx, as.ArmamentQuery, as, as.inters, v)
}

func (as *ArmamentSelect) sqlScan(ctx context.Context, root *ArmamentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}