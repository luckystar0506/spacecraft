// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"spacecraft/ent/armament"
	"spacecraft/ent/predicate"
	"spacecraft/ent/spacecraft"
	"spacecraft/ent/spacecraftarmament"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpacecraftArmamentQuery is the builder for querying SpacecraftArmament entities.
type SpacecraftArmamentQuery struct {
	config
	ctx            *QueryContext
	order          []spacecraftarmament.OrderOption
	inters         []Interceptor
	predicates     []predicate.SpacecraftArmament
	withSpacecraft *SpacecraftQuery
	withArmament   *ArmamentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SpacecraftArmamentQuery builder.
func (saq *SpacecraftArmamentQuery) Where(ps ...predicate.SpacecraftArmament) *SpacecraftArmamentQuery {
	saq.predicates = append(saq.predicates, ps...)
	return saq
}

// Limit the number of records to be returned by this query.
func (saq *SpacecraftArmamentQuery) Limit(limit int) *SpacecraftArmamentQuery {
	saq.ctx.Limit = &limit
	return saq
}

// Offset to start from.
func (saq *SpacecraftArmamentQuery) Offset(offset int) *SpacecraftArmamentQuery {
	saq.ctx.Offset = &offset
	return saq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (saq *SpacecraftArmamentQuery) Unique(unique bool) *SpacecraftArmamentQuery {
	saq.ctx.Unique = &unique
	return saq
}

// Order specifies how the records should be ordered.
func (saq *SpacecraftArmamentQuery) Order(o ...spacecraftarmament.OrderOption) *SpacecraftArmamentQuery {
	saq.order = append(saq.order, o...)
	return saq
}

// QuerySpacecraft chains the current query on the "spacecraft" edge.
func (saq *SpacecraftArmamentQuery) QuerySpacecraft() *SpacecraftQuery {
	query := (&SpacecraftClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(spacecraftarmament.Table, spacecraftarmament.FieldID, selector),
			sqlgraph.To(spacecraft.Table, spacecraft.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, spacecraftarmament.SpacecraftTable, spacecraftarmament.SpacecraftColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryArmament chains the current query on the "armament" edge.
func (saq *SpacecraftArmamentQuery) QueryArmament() *ArmamentQuery {
	query := (&ArmamentClient{config: saq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := saq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := saq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(spacecraftarmament.Table, spacecraftarmament.FieldID, selector),
			sqlgraph.To(armament.Table, armament.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, spacecraftarmament.ArmamentTable, spacecraftarmament.ArmamentColumn),
		)
		fromU = sqlgraph.SetNeighbors(saq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SpacecraftArmament entity from the query.
// Returns a *NotFoundError when no SpacecraftArmament was found.
func (saq *SpacecraftArmamentQuery) First(ctx context.Context) (*SpacecraftArmament, error) {
	nodes, err := saq.Limit(1).All(setContextOp(ctx, saq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{spacecraftarmament.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) FirstX(ctx context.Context) *SpacecraftArmament {
	node, err := saq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SpacecraftArmament ID from the query.
// Returns a *NotFoundError when no SpacecraftArmament ID was found.
func (saq *SpacecraftArmamentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(1).IDs(setContextOp(ctx, saq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{spacecraftarmament.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) FirstIDX(ctx context.Context) int {
	id, err := saq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SpacecraftArmament entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SpacecraftArmament entity is found.
// Returns a *NotFoundError when no SpacecraftArmament entities are found.
func (saq *SpacecraftArmamentQuery) Only(ctx context.Context) (*SpacecraftArmament, error) {
	nodes, err := saq.Limit(2).All(setContextOp(ctx, saq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{spacecraftarmament.Label}
	default:
		return nil, &NotSingularError{spacecraftarmament.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) OnlyX(ctx context.Context) *SpacecraftArmament {
	node, err := saq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SpacecraftArmament ID in the query.
// Returns a *NotSingularError when more than one SpacecraftArmament ID is found.
// Returns a *NotFoundError when no entities are found.
func (saq *SpacecraftArmamentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = saq.Limit(2).IDs(setContextOp(ctx, saq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{spacecraftarmament.Label}
	default:
		err = &NotSingularError{spacecraftarmament.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) OnlyIDX(ctx context.Context) int {
	id, err := saq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SpacecraftArmaments.
func (saq *SpacecraftArmamentQuery) All(ctx context.Context) ([]*SpacecraftArmament, error) {
	ctx = setContextOp(ctx, saq.ctx, "All")
	if err := saq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SpacecraftArmament, *SpacecraftArmamentQuery]()
	return withInterceptors[[]*SpacecraftArmament](ctx, saq, qr, saq.inters)
}

// AllX is like All, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) AllX(ctx context.Context) []*SpacecraftArmament {
	nodes, err := saq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SpacecraftArmament IDs.
func (saq *SpacecraftArmamentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if saq.ctx.Unique == nil && saq.path != nil {
		saq.Unique(true)
	}
	ctx = setContextOp(ctx, saq.ctx, "IDs")
	if err = saq.Select(spacecraftarmament.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) IDsX(ctx context.Context) []int {
	ids, err := saq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (saq *SpacecraftArmamentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, saq.ctx, "Count")
	if err := saq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, saq, querierCount[*SpacecraftArmamentQuery](), saq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) CountX(ctx context.Context) int {
	count, err := saq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (saq *SpacecraftArmamentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, saq.ctx, "Exist")
	switch _, err := saq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (saq *SpacecraftArmamentQuery) ExistX(ctx context.Context) bool {
	exist, err := saq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SpacecraftArmamentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (saq *SpacecraftArmamentQuery) Clone() *SpacecraftArmamentQuery {
	if saq == nil {
		return nil
	}
	return &SpacecraftArmamentQuery{
		config:         saq.config,
		ctx:            saq.ctx.Clone(),
		order:          append([]spacecraftarmament.OrderOption{}, saq.order...),
		inters:         append([]Interceptor{}, saq.inters...),
		predicates:     append([]predicate.SpacecraftArmament{}, saq.predicates...),
		withSpacecraft: saq.withSpacecraft.Clone(),
		withArmament:   saq.withArmament.Clone(),
		// clone intermediate query.
		sql:  saq.sql.Clone(),
		path: saq.path,
	}
}

// WithSpacecraft tells the query-builder to eager-load the nodes that are connected to
// the "spacecraft" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SpacecraftArmamentQuery) WithSpacecraft(opts ...func(*SpacecraftQuery)) *SpacecraftArmamentQuery {
	query := (&SpacecraftClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withSpacecraft = query
	return saq
}

// WithArmament tells the query-builder to eager-load the nodes that are connected to
// the "armament" edge. The optional arguments are used to configure the query builder of the edge.
func (saq *SpacecraftArmamentQuery) WithArmament(opts ...func(*ArmamentQuery)) *SpacecraftArmamentQuery {
	query := (&ArmamentClient{config: saq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	saq.withArmament = query
	return saq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Qty int `json:"qty,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SpacecraftArmament.Query().
//		GroupBy(spacecraftarmament.FieldQty).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (saq *SpacecraftArmamentQuery) GroupBy(field string, fields ...string) *SpacecraftArmamentGroupBy {
	saq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SpacecraftArmamentGroupBy{build: saq}
	grbuild.flds = &saq.ctx.Fields
	grbuild.label = spacecraftarmament.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Qty int `json:"qty,omitempty"`
//	}
//
//	client.SpacecraftArmament.Query().
//		Select(spacecraftarmament.FieldQty).
//		Scan(ctx, &v)
func (saq *SpacecraftArmamentQuery) Select(fields ...string) *SpacecraftArmamentSelect {
	saq.ctx.Fields = append(saq.ctx.Fields, fields...)
	sbuild := &SpacecraftArmamentSelect{SpacecraftArmamentQuery: saq}
	sbuild.label = spacecraftarmament.Label
	sbuild.flds, sbuild.scan = &saq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SpacecraftArmamentSelect configured with the given aggregations.
func (saq *SpacecraftArmamentQuery) Aggregate(fns ...AggregateFunc) *SpacecraftArmamentSelect {
	return saq.Select().Aggregate(fns...)
}

func (saq *SpacecraftArmamentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range saq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, saq); err != nil {
				return err
			}
		}
	}
	for _, f := range saq.ctx.Fields {
		if !spacecraftarmament.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if saq.path != nil {
		prev, err := saq.path(ctx)
		if err != nil {
			return err
		}
		saq.sql = prev
	}
	return nil
}

func (saq *SpacecraftArmamentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SpacecraftArmament, error) {
	var (
		nodes       = []*SpacecraftArmament{}
		_spec       = saq.querySpec()
		loadedTypes = [2]bool{
			saq.withSpacecraft != nil,
			saq.withArmament != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SpacecraftArmament).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SpacecraftArmament{config: saq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, saq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := saq.withSpacecraft; query != nil {
		if err := saq.loadSpacecraft(ctx, query, nodes, nil,
			func(n *SpacecraftArmament, e *Spacecraft) { n.Edges.Spacecraft = e }); err != nil {
			return nil, err
		}
	}
	if query := saq.withArmament; query != nil {
		if err := saq.loadArmament(ctx, query, nodes, nil,
			func(n *SpacecraftArmament, e *Armament) { n.Edges.Armament = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (saq *SpacecraftArmamentQuery) loadSpacecraft(ctx context.Context, query *SpacecraftQuery, nodes []*SpacecraftArmament, init func(*SpacecraftArmament), assign func(*SpacecraftArmament, *Spacecraft)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SpacecraftArmament)
	for i := range nodes {
		fk := nodes[i].SpacecraftID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(spacecraft.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "spacecraft_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (saq *SpacecraftArmamentQuery) loadArmament(ctx context.Context, query *ArmamentQuery, nodes []*SpacecraftArmament, init func(*SpacecraftArmament), assign func(*SpacecraftArmament, *Armament)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SpacecraftArmament)
	for i := range nodes {
		fk := nodes[i].ArmamentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(armament.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "armament_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (saq *SpacecraftArmamentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := saq.querySpec()
	_spec.Node.Columns = saq.ctx.Fields
	if len(saq.ctx.Fields) > 0 {
		_spec.Unique = saq.ctx.Unique != nil && *saq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, saq.driver, _spec)
}

func (saq *SpacecraftArmamentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(spacecraftarmament.Table, spacecraftarmament.Columns, sqlgraph.NewFieldSpec(spacecraftarmament.FieldID, field.TypeInt))
	_spec.From = saq.sql
	if unique := saq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if saq.path != nil {
		_spec.Unique = true
	}
	if fields := saq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, spacecraftarmament.FieldID)
		for i := range fields {
			if fields[i] != spacecraftarmament.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if saq.withSpacecraft != nil {
			_spec.Node.AddColumnOnce(spacecraftarmament.FieldSpacecraftID)
		}
		if saq.withArmament != nil {
			_spec.Node.AddColumnOnce(spacecraftarmament.FieldArmamentID)
		}
	}
	if ps := saq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := saq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := saq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := saq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (saq *SpacecraftArmamentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(saq.driver.Dialect())
	t1 := builder.Table(spacecraftarmament.Table)
	columns := saq.ctx.Fields
	if len(columns) == 0 {
		columns = spacecraftarmament.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if saq.sql != nil {
		selector = saq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if saq.ctx.Unique != nil && *saq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range saq.predicates {
		p(selector)
	}
	for _, p := range saq.order {
		p(selector)
	}
	if offset := saq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := saq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SpacecraftArmamentGroupBy is the group-by builder for SpacecraftArmament entities.
type SpacecraftArmamentGroupBy struct {
	selector
	build *SpacecraftArmamentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sagb *SpacecraftArmamentGroupBy) Aggregate(fns ...AggregateFunc) *SpacecraftArmamentGroupBy {
	sagb.fns = append(sagb.fns, fns...)
	return sagb
}

// Scan applies the selector query and scans the result into the given value.
func (sagb *SpacecraftArmamentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sagb.build.ctx, "GroupBy")
	if err := sagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpacecraftArmamentQuery, *SpacecraftArmamentGroupBy](ctx, sagb.build, sagb, sagb.build.inters, v)
}

func (sagb *SpacecraftArmamentGroupBy) sqlScan(ctx context.Context, root *SpacecraftArmamentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sagb.fns))
	for _, fn := range sagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sagb.flds)+len(sagb.fns))
		for _, f := range *sagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SpacecraftArmamentSelect is the builder for selecting fields of SpacecraftArmament entities.
type SpacecraftArmamentSelect struct {
	*SpacecraftArmamentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sas *SpacecraftArmamentSelect) Aggregate(fns ...AggregateFunc) *SpacecraftArmamentSelect {
	sas.fns = append(sas.fns, fns...)
	return sas
}

// Scan applies the selector query and scans the result into the given value.
func (sas *SpacecraftArmamentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sas.ctx, "Select")
	if err := sas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SpacecraftArmamentQuery, *SpacecraftArmamentSelect](ctx, sas.SpacecraftArmamentQuery, sas, sas.inters, v)
}

func (sas *SpacecraftArmamentSelect) sqlScan(ctx context.Context, root *SpacecraftArmamentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sas.fns))
	for _, fn := range sas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
